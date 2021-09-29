package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/pkg/utils"
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
			type Config struct {
				Email                       string `mapstructure:"email" validate:"required"`
				Password                    string `mapstructure:"password" validate:"required_unless=PasswordStdin false"`
				PasswordStdin               bool   `mapstructure:"password_stdin"`
				cmdCommon.MySQLConfigCommon `mapstructure:",squash"`
			}

			var config Config

			if err := viper.Unmarshal(&config); err != nil {
				log.Fatalf("Unable to unmarshal configuration: %s", err)
			}

			// NOTE we don't support passing email, password, password-stdin via env, so manually extract and apply
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			passwordStdin, _ := cmd.Flags().GetBool("password-stdin")

			config.Email = email
			config.Password = password
			config.PasswordStdin = passwordStdin

			v := validator.New()
			err := v.Struct(config)

			if err != nil {
				log.Fatalf("Configuration is invalid: %s", err)
			}

			logger, err := zap.NewDevelopment()

			if err != nil {
				log.Fatalf("Unable to create a logger: %s", err)
			}

			logger.Debug("BRANGUS")
		},
	}
)

func init() {
	flags := loginCommand.Flags()

	cmdCommon.MySQLFlags(flags)

	flags.StringP("email", "e", "", "The email to log in with.")
	flags.StringP("password", "p", "", "The password to log in with.")
	flags.BoolP("password-stdin", "", false, "Read the password from standard input.")

	utils.Nop()
}
