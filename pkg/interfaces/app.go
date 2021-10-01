package interfaces

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App interface {
	Config() AppConfig
	Dao() AppDaos
	DB() *gorm.DB
	Logger() *zap.Logger
	Service() AppServices
	Router() *mux.Router
	RootLogger() *zap.Logger
}

type AppServices interface {
	Login() LoginService
	JWT() JWTService
}

type AppDaos interface {
	Users() UserDao
}
