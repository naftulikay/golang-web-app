package serve

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/naftulikay/golang-webapp/pkg"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func serveCommandFlagsToEnv() map[string]string {
	return map[string]string{
		cmdConstants.CliFlagMySQLHost:     cmdConstants.EnvVarMySQLHost,
		cmdConstants.CliFlagMySQLPort:     cmdConstants.EnvVarMySQLPort,
		cmdConstants.CliFlagMySQLDatabase: cmdConstants.EnvVarMySQLDatabase,
		cmdConstants.CliFlagMySQLUser:     cmdConstants.EnvVarMySQLUser,
		cmdConstants.CliFlagMySQLPassword: cmdConstants.EnvVarMySQLPassword,
	}
}

func serveCommandEnvDefaults() map[string]interface{} {
	return map[string]interface{}{
		cmdConstants.EnvVarMySQLHost:  cmdConstants.DefaultMySQLHost,
		cmdConstants.EnvVarMySQLPort:  cmdConstants.DefaultMySQLPort,
		cmdConstants.EnvVarListenHost: cmdConstants.DefaultListenHost,
		cmdConstants.EnvVarListenPort: cmdConstants.DefaultListenPort,
	}
}

type ServeConfig struct { // implements interfaces.AppConfig
	Environment                 string `mapstructure:"env"`
	ListenHost                  string `mapstructure:"listen_host"`
	ListenPort                  uint16 `mapstructure:"listen_port"`
	cmdCommon.MySQLConfigCommon `mapstructure:",squash"`
}

func (s ServeConfig) Env() string {
	return s.Environment
}

func (s ServeConfig) Listen() interfaces.ListenConfig {
	return ServeListenConfig{
		listenHost: s.ListenHost,
		listenPort: s.ListenPort,
	}
}

func (s ServeConfig) MySQL() interfaces.MySQLConfig {
	return ServeMySQLConfig{
		MySQLConfigCommon: cmdCommon.MySQLConfigCommon{
			MySQLHost:     s.MySQLHost,
			MySQLPort:     s.MySQLPort,
			MySQLDatabase: s.MySQLDatabase,
			MySQLUser:     s.MySQLUser,
			MySQLPassword: s.MySQLPassword,
		},
	}
}

type ServeListenConfig struct { // implements interfaces.ListenConfig
	listenHost string `mapstructure:"listen_host"`
	listenPort uint16 `mapstructure:"listen_port"`
}

func (s ServeListenConfig) Host() string {
	return s.listenHost
}

func (s ServeListenConfig) Port() uint16 {
	return s.listenPort
}

type ServeMySQLConfig struct { // implements interfaces.MySQLConfig
	cmdCommon.MySQLConfigCommon `mapstructure:",squash"`
}

func (s ServeMySQLConfig) Host() string {
	return s.MySQLHost
}

func (s ServeMySQLConfig) Port() uint16 {
	return s.MySQLPort
}

func (s ServeMySQLConfig) Database() string {
	return s.MySQLDatabase
}

func (s ServeMySQLConfig) User() string {
	return s.MySQLUser
}

func (s ServeMySQLConfig) Password() string {
	return s.MySQLPassword
}

var (
	serveCommand = &cobra.Command{
		Use:   "serve",
		Short: "Run the web service.",
		PreRun: func(cmd *cobra.Command, args []string) {
			for flag, env := range serveCommandFlagsToEnv() {
				if err := viper.BindPFlag(env, cmd.Flags().Lookup(flag)); err != nil {
					log.Fatalf("Unable to bind environment variable %s to CLI flag %s: %s", strings.ToUpper(env),
						fmt.Sprintf("--%s", flag), err)
				}
			}

			for env, defaultValue := range serveCommandEnvDefaults() {
				viper.SetDefault(env, defaultValue)
			}

			// NOTE if you need to add a viper config file, register it here
		},
		Run: func(cmd *cobra.Command, args []string) {
			var config ServeConfig

			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Unable to unmarshal config: %s", err)
			}

			// validation for non-default parameters
			if len(config.ListenHost) == 0 {
				log.Fatalf("Please pass a valid listen host via the --%s CLI flag or %s environment variable.",
					cmdConstants.CliFlagListen, strings.ToUpper(cmdConstants.EnvVarListenHost))
			}

			if len(config.MySQLDatabase) == 0 {
				log.Fatalf("Please pass a MySQL database schema name via the --%s CLI flag or %s environment variable.",
					cmdConstants.CliFlagMySQLDatabase, strings.ToUpper(cmdConstants.EnvVarMySQLDatabase))
			}

			if len(config.MySQLUser) == 0 {
				log.Fatalf("Please pass a MySQL user name via the --%s CLI flag or %s environment variable.",
					cmdConstants.CliFlagMySQLUser, strings.ToUpper(cmdConstants.EnvVarMySQLUser))
			}

			if len(config.MySQLPassword) == 0 {
				log.Fatalf("Please pass a MySQL password via the --%s CLI flag or %s environment variable.",
					cmdConstants.CliFlagMySQLPassword, strings.ToUpper(cmdConstants.EnvVarMySQLPassword))
			}

			var db *gorm.DB

			// connect to database
			err := backoff.Retry(func() error {
				uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", config.MySQLUser, config.MySQLPassword,
					config.MySQLHost, config.MySQLPort, config.MySQLDatabase)

				connection, err := gorm.Open(mysql.Open(uri))

				if err != nil {
					return err
				}

				db = connection

				return nil
			}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 15))

			if err != nil {
				log.Fatalf("Unable to connect to database: %s", err)
			}

			log.Printf("Database connected (%v)", db)
			log.Printf("Initialization complete, serving at http://%s:%d/", config.ListenHost, config.ListenPort)

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/get", func(w http.ResponseWriter, r *http.Request) {
				log.Printf("Received request!")
				w.WriteHeader(http.StatusOK)
			})

			pkg.Start(config)

			server := &http.Server{
				Addr:         fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort),
				Handler:      mux,
				ReadTimeout:  300 * time.Second, // support long debugging sessions
				WriteTimeout: 300 * time.Second,
			}

			err = server.ListenAndServe()

			if err != nil {
				log.Fatalf("Failed to serve application: %s", err)
			}
		},
	}
)

func Commands() []*cobra.Command {
	return []*cobra.Command{serveCommand}
}

func init() {
	flags := serveCommand.Flags()

	// --mysql-host
	flags.StringP(cmdConstants.CliFlagMySQLHost, "", cmdConstants.DefaultMySQLHost,
		fmt.Sprintf("The hostname or IP address of the MySQL database server. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarMySQLHost)))
	// --mysql-port
	flags.Uint16P(cmdConstants.CliFlagMySQLPort, "", cmdConstants.DefaultMySQLPort,
		fmt.Sprintf("The port which the MySQL database server is listening on. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarMySQLPort)))
	// --mysql-database
	flags.StringP(cmdConstants.CliFlagMySQLDatabase, "", "",
		fmt.Sprintf("The database schema within MySQL to use for the application. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarMySQLDatabase)))
	// --mysql-user
	flags.StringP(cmdConstants.CliFlagMySQLUser, "", "",
		fmt.Sprintf("The MySQL user to connect as. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarMySQLUser)))
	// --mysql-port
	flags.StringP(cmdConstants.CliFlagMySQLPassword, "", "",
		fmt.Sprintf("The MySQL password to connect with. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarMySQLPassword)))
	// --listen
	flags.StringP(cmdConstants.CliFlagListen, "H", cmdConstants.DefaultListenHost,
		fmt.Sprintf("The host to listen on for incoming connections. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarListenHost)))
	// --port
	flags.Uint16P(cmdConstants.CliFlagPort, "p", cmdConstants.DefaultListenPort,
		fmt.Sprintf("The port to listen on for incoming connections. [env: %s]",
			strings.ToUpper(cmdConstants.CliFlagPort)))
}
