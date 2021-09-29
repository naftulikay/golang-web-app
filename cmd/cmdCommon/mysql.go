package cmdCommon

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/naftulikay/golang-webapp/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

func MySQLEnvToCli() map[string]string {
	return map[string]string{
		cmdConstants.EnvVarMySQLHost:     cmdConstants.CliFlagMySQLHost,
		cmdConstants.EnvVarMySQLPort:     cmdConstants.CliFlagMySQLPort,
		cmdConstants.EnvVarMySQLDatabase: cmdConstants.CliFlagMySQLDatabase,
		cmdConstants.EnvVarMySQLUser:     cmdConstants.CliFlagMySQLUser,
		cmdConstants.EnvVarMySQLPassword: cmdConstants.CliFlagMySQLPassword,
	}
}

func MySQLEnvDefaults() map[string]interface{} {
	return map[string]interface{}{
		cmdConstants.EnvVarMySQLHost: cmdConstants.DefaultMySQLHost,
		cmdConstants.EnvVarMySQLPort: cmdConstants.DefaultMySQLPort,
	}
}

func MySQLFlags(flags *pflag.FlagSet) {
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

	utils.Nop() // debug point
}

func MySQLBindFlagsToEnv(flags *pflag.FlagSet) error {
	var result error

	for env, flag := range MySQLEnvToCli() {
		if err := viper.BindPFlag(env, flags.Lookup(flag)); err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result
}

func MySQLValidate(config MySQLConfigCommon) error {
	var result error

	v := validator.New()
	err := v.Struct(config)

	if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
		return invalidErr
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, validationErr := range validationErrors {
			result = multierror.Append(result, validationErr)
		}
	}

	return result
}

func MySQLRegisterDefaults() {
	for env, def := range MySQLEnvDefaults() {
		viper.SetDefault(env, def)
	}
}
