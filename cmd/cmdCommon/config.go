package cmdCommon

type MySQLConfigCommon struct {
	MySQLHost     string `mapstructure:"mysql_host" govalid:"req"`
	MySQLPort     uint16 `mapstructure:"mysql_port" govalid:"req"`
	MySQLDatabase string `mapstructure:"mysql_database" govalid:"req"`
	MySQLUser     string `mapstructure:"mysql_user" govalid:"req"`
	MySQLPassword string `mapstructure:"mysql_password" govalid:"req"`
}
