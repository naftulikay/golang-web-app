package serve

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/naftulikay/golang-webapp/pkg"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twharmon/govalid"
	"log"
	"strings"
)

func serveCommandFlagsToEnv() map[string]string {
	return map[string]string{
		cmdConstants.CliFlagEnv:           cmdConstants.EnvVarEnvironment,
		cmdConstants.CliFlagMySQLHost:     cmdConstants.EnvVarMySQLHost,
		cmdConstants.CliFlagMySQLPort:     cmdConstants.EnvVarMySQLPort,
		cmdConstants.CliFlagMySQLDatabase: cmdConstants.EnvVarMySQLDatabase,
		cmdConstants.CliFlagMySQLUser:     cmdConstants.EnvVarMySQLUser,
		cmdConstants.CliFlagMySQLPassword: cmdConstants.EnvVarMySQLPassword,
		cmdConstants.CliFlagListen:        cmdConstants.EnvVarListenHost,
		cmdConstants.CliFlagPort:          cmdConstants.EnvVarListenPort,
	}
}

func serveCommandEnvDefaults() map[string]interface{} {
	return map[string]interface{}{
		cmdConstants.EnvVarMySQLHost:   cmdConstants.DefaultMySQLHost,
		cmdConstants.EnvVarMySQLPort:   cmdConstants.DefaultMySQLPort,
		cmdConstants.EnvVarListenHost:  cmdConstants.DefaultListenHost,
		cmdConstants.EnvVarListenPort:  cmdConstants.DefaultListenPort,
		cmdConstants.EnvVarEnvironment: cmdConstants.DefaultEnvironment,
	}
}

type ServeConfig struct { // implements interfaces.AppConfig
	Environment                 string `mapstructure:"env" govalid:"req|in:dev,stg,prod"`
	ListenHost                  string `mapstructure:"listen_host" govalid:"req"`
	ListenPort                  uint16 `mapstructure:"listen_port" govalid:"req"`
	cmdCommon.MySQLConfigCommon `mapstructure:",squash" govalid:"req"`
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
	validator    *govalid.Validator
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

			violations, err := validator.Violations(config)

			if err != nil {
				log.Fatalf("Unable to run validator: %s", err)
			}

			if len(violations) > 0 {
				for _, v := range violations {
					log.Printf("Configuration Error: %s", v)
				}

				log.Fatalf("Validation failed.")
			}

			violations, err = validator.Violations(config.MySQLConfigCommon)

			if err != nil {
				log.Fatalf("Unable to run validator: %s", err)
			}

			if len(violations) > 0 {
				for _, v := range violations {
					log.Printf("MySQL Configuration Error: %s", v)
				}

				log.Fatalf("Validation failed.")
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

			pkg.Start(config)
		},
	}
)

func Commands() []*cobra.Command {
	return []*cobra.Command{serveCommand}
}

func init() {
	validator = govalid.New()

	err := validator.Register(ServeConfig{}, ServeListenConfig{}, ServeMySQLConfig{},
		cmdCommon.MySQLConfigCommon{})

	if err != nil {
		log.Fatalf("Unable to configure validators: %s", err)
	}

	flags := serveCommand.Flags()

	// --env
	flags.StringP(cmdConstants.CliFlagEnv, "e", cmdConstants.DefaultEnvironment,
		fmt.Sprintf("The execution environment for the application. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarEnvironment)))

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
