package cmdConstants

const (
	EnvVarMySQLHost     = "mysql_host"
	EnvVarMySQLPort     = "mysql_port"
	EnvVarMySQLDatabase = "mysql_database"
	EnvVarMySQLUser     = "mysql_user"
	EnvVarMySQLPassword = "mysql_password"
	EnvVarListenHost    = "listen_host"
	EnvVarListenPort    = "listen_port"
)

// EnvVariables Return a list of all known environment variables used by any and all commands.
func EnvVariables() []string {
	return []string{
		EnvVarMySQLHost, EnvVarMySQLPort, EnvVarMySQLDatabase, EnvVarMySQLUser, EnvVarMySQLPassword,
	}
}
