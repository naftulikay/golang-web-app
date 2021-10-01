package cmdCommon

type MySQLConfigCommon struct {
	MySQLHost     string `mapstructure:"mysql_host" govalid:"req" validate:"required,hostname|ip"`
	MySQLPort     uint16 `mapstructure:"mysql_port" govalid:"req" validate:"required,gt=0"`
	MySQLDatabase string `mapstructure:"mysql_database" govalid:"req" validate:"required"`
	MySQLUser     string `mapstructure:"mysql_user" govalid:"req" validate:"required"`
	MySQLPassword string `mapstructure:"mysql_password" govalid:"req" validate:"required"`
}

func (m MySQLConfigCommon) Host() string {
	return m.MySQLHost
}

func (m MySQLConfigCommon) Port() uint16 {
	return m.MySQLPort
}

func (m MySQLConfigCommon) Database() string {
	return m.MySQLDatabase
}

func (m MySQLConfigCommon) User() string {
	return m.MySQLUser
}

func (m MySQLConfigCommon) Password() string {
	return m.MySQLPassword
}
