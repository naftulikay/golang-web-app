package service

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdCommon"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
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

			if err := viper.BindPFlag(cmdConstants.EnvVarJWTKey, cmd.Flags().Lookup(cmdConstants.CliFlagJWTKey)); err != nil {
				log.Fatalf("Unable to bind JWT key variable to CLI flags: %s", err)
			}

			cmdCommon.MySQLRegisterDefaults()
		},
		Run: func(cmd *cobra.Command, args []string) {
			root := cmdCommon.Logger()
			logger := root.Named("cmd.services.login")
			appLogger := root.Named("app")

			type Config struct {
				Email                       string `mapstructure:"email" validate:"required,email"`
				Password                    string `mapstructure:"password" validate:"required_unless=PasswordStdin true"`
				PasswordStdin               bool   `mapstructure:"password_stdin"`
				JWTKey                      string `mapstructure:"jwt_key" validate:"required,base64|hexadecimal"`
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

			// parse JWT key
			jwtKeyBytes, err := cmdCommon.HexOrBase64(config.JWTKey)

			if err != nil {
				logger.Fatal("Unable to decode JWT key as base-64 or hex", zap.Error(err))
			}

			if len(jwtKeyBytes) != 32 {
				logger.Fatal("JWT key is not 32 bytes long.", zap.Int("length", len(jwtKeyBytes)))
			}

			var jwtKey [32]byte

			copy(jwtKey[:], jwtKeyBytes)

			loginSvc, err := initializeLoginService(
				config.MySQLConfigCommon,
				jwtKey,
				appLogger.Named("services.login"),
				appLogger.Named("services.jwt"),
				appLogger.Named("dao.user"),
				appLogger.Named("dbinit"),
			)

			if err != nil {
				logger.Fatal("Unable to wire dependencies.", zap.Error(err))
			}

			// ready to go
			logger.Info(fmt.Sprintf("Attempting to log in user %s...", config.Email))

			login, err := loginSvc.Login(config.Email, config.Password)

			if err != nil {
				logger.Fatal("Login failed.")
			}

			user := (*login).User()

			logger.Info("Login successful", zap.Uint("user_id", (*login).User().ID),
				zap.String("user_email", user.Email),
				zap.String("user_name", user.Name()),
				zap.String("jwt_token", (*login).SignedToken()))
		},
	}
)

func init() {
	flags := loginCommand.Flags()

	cmdCommon.MySQLFlags(flags)

	flags.StringP("email", "e", "", "The email to log in with.")
	flags.StringP("password", "p", "", "The password to log in with.")
	flags.BoolP("password-stdin", "", false, "Read the password from standard input.")

	flags.StringP(cmdConstants.CliFlagJWTKey, "j", "",
		fmt.Sprintf("The 32-byte JWT key encoded in hex or base64. [env: %s]", cmdConstants.EnvVarJWTKey))
}
