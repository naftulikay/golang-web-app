package serve

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/naftulikay/golang-webapp/pkg"
	"github.com/naftulikay/golang-webapp/pkg/constants"
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
		cmdConstants.CliFlagJWTKey:        cmdConstants.EnvVarJWTKey,
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
	JWTKey                      string `mapstructure:"jwt_key" validate:"required,base64|hexadecimal"`
	Migrate                     bool   `mapstructure:"migrate"`
	Environment                 string `mapstructure:"env" validate:"required,oneof=dev stg prod"`
	ListenHost                  string `mapstructure:"listen_host" validate:"required"`
	ListenPort                  uint16 `mapstructure:"listen_port" validate:"required,gt=0"`
	cmdCommon.MySQLConfigCommon `mapstructure:",squash" govalid:"req"`
}

func (s ServeConfig) AutoMigrate() bool {
	return s.Migrate
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

			migrate, err := cmd.Flags().GetBool(cmdConstants.CliFlagMigrate)

			if err == nil {
				config.Migrate = migrate
			}

			cmdCommon.ValidateConfig(config, nil)

			jwtKeyBytes, err := cmdCommon.HexOrBase64(config.JWTKey)

			if err != nil {
				log.Fatalf("Unable to parse JWT key as either hexadecimal or base-64: %s", err)
			}

			if len(jwtKeyBytes) != 32 {
				log.Fatalf("JWT key, when decoded, must be 32 bytes in length, not %d", len(jwtKeyBytes))
			}

			var jwtKey [32]byte

			copy(jwtKey[:], jwtKeyBytes)

			pkg.Start(config, jwtKey)
		},
	}
)

func Commands() []*cobra.Command {
	return []*cobra.Command{serveCommand}
}

func init() {
	flags := serveCommand.Flags()

	cmdCommon.MySQLFlags(flags)
	cmdCommon.ListenFlags(flags)

	// --env
	flags.StringP(cmdConstants.CliFlagEnv, "", constants.DevEnvironment,
		fmt.Sprintf("The runtime execution environment. [env: %s]", strings.ToUpper(cmdConstants.EnvVarEnvironment)))

	flags.StringP(cmdConstants.CliFlagJWTKey, "j", "",
		fmt.Sprintf("The JWT key, which is 32 bytes of binary data encoded either as hex or as base-64. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarJWTKey)))

	// --migrate
	flags.BoolP(cmdConstants.CliFlagMigrate, "", false,
		"Whether to attempt to auto-migrate database models on start.")
}
