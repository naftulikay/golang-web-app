package serve

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
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
		},
		Run: func(cmd *cobra.Command, args []string) {
			type Config struct {
				ListenHost                  string `mapstructure:"listen_host"`
				ListenPort                  uint16 `mapstructure:"listen_port"'`
				cmdCommon.MySQLConfigCommon `mapstructure:",squash"`
			}

			var config Config

			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Unable to unmarshal config: %s", err)
			}

			log.Printf("Config: %+v", config)
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
