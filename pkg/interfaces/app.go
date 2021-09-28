package interfaces

type App interface {
	Config() AppConfig
	Dao() AppDaos
	Service() AppServices
}

type AppServices interface {
	Login() LoginService
}

type AppDaos interface {
	Users() UserDao
}
