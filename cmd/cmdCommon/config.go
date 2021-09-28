package cmdCommon

type MySQLConfigCommon struct {
	MySQLHost     string `mapstructure:"mysql_host"`
	MySQLPort     uint16 `mapstructure:"mysql_port"`
	MySQLDatabase string `mapstructure:"mysql_database"`
	MySQLUser     string `mapstructure:"mysql_user"`
	MySQLPassword string `mapstructure:"mysql_password"`
}
