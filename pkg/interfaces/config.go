package interfaces

type AppConfig interface {
	Env() string
	Listen() ListenConfig
	MySQL() MySQLConfig
}

type MySQLConfig interface {
	Host() string
	Port() uint16
	Database() string
	User() string
	Password() string
}

type ListenConfig interface {
	Host() string
	Port() uint16
}
