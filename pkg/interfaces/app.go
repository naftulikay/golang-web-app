package interfaces

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App interface {
	Config() *AppConfig
	Dao() *AppDaos
	DB() *gorm.DB
	Service() *AppServices
	Router() *mux.Router
	RootLogger() *zap.Logger
}

type AppServices interface {
	Login() *LoginService
}

type AppDaos interface {
	Users() *UserDao
}
