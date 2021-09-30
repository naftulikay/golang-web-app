package service

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

var (
	loginCommand = &cobra.Command{
		Use:   "login",
		Short: "Use the login service to test login and token generation.",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := cmdCommon.MySQLBindFlagsToEnv(cmd.Flags()); err != nil {
				log.Fatalf("Unable to bind MySQL environment variable to CLI flags: %s", err)
			}

			cmdCommon.MySQLRegisterDefaults()
		},
		Run: func(cmd *cobra.Command, args []string) {
			root := cmdCommon.Logger()
			logger := root.Named("cmd.services.login")

			type Config struct {
				Email                       string `mapstructure:"email" validate:"required,email"`
				Password                    string `mapstructure:"password" validate:"required_unless=PasswordStdin true"`
				PasswordStdin               bool   `mapstructure:"password_stdin"`
				cmdCommon.MySQLConfigCommon `mapstructure:",squash"`
			}

			var config Config

			if err := viper.Unmarshal(&config); err != nil {
				logger.Fatal("Unable to unmarshal configuration", zap.Error(err))
			}

			// NOTE we don't support passing email, password, password-stdin via env, so manually extract and apply
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			passwordStdin, _ := cmd.Flags().GetBool("password-stdin")

			config.Email = email
			config.Password = password
			config.PasswordStdin = passwordStdin

			cmdCommon.ValidateConfig(config, logger)

			if passwordStdin {
				config.Password = cmdCommon.StdinPassword("Enter Password: ", logger)
			}

			logger.Info(fmt.Sprintf("Attempting to log in user %s...", config.Email))
		},
	}
)

func LoggerBuilder(parent *zap.Logger, name string) func() *zap.Logger {
	return func() *zap.Logger {
		return parent.Named(name)
	}
}

func init() {
	flags := loginCommand.Flags()

	cmdCommon.MySQLFlags(flags)

	flags.StringP("email", "e", "", "The email to log in with.")
	flags.StringP("password", "p", "", "The password to log in with.")
	flags.BoolP("password-stdin", "", false, "Read the password from standard input.")
}
