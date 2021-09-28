package service

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AppImpl Implementation of the interfaces.App for single point of entry to application.
type AppImpl struct {
	db         *gorm.DB
	config     *interfaces.AppConfig
	services   *interfaces.AppServices
	daos       *interfaces.AppDaos
	router     *mux.Router
	rootLogger *zap.Logger
}

func (a AppImpl) DB() *gorm.DB {
	return a.db
}

func (a AppImpl) Config() *interfaces.AppConfig {
	return a.config
}

func (a AppImpl) Dao() *interfaces.AppDaos {
	return a.daos
}

func (a AppImpl) Service() *interfaces.AppServices {
	return a.services
}

func (a AppImpl) Router() *mux.Router {
	return a.router
}

func (a AppImpl) RootLogger() *zap.Logger {
	return a.rootLogger
}

// AppServicesImpl Implementation of the interfaces.AppServices interface for exposing services to the application.
type AppServicesImpl struct {
	login *interfaces.LoginService
}

func (a AppServicesImpl) Login() *interfaces.LoginService {
	return a.login
}

// AppDaosImpl Implementation of the interfaces.AppDaos interface for exposing DAOs to the application.
type AppDaosImpl struct {
	users *interfaces.UserDao
}

func (a AppDaosImpl) Users() *interfaces.UserDao {
	return a.users
}
