package service

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AppImpl Implementation of the interfaces.App for single point of entry to application.
type AppImpl struct {
	db         *gorm.DB
	daos       interfaces.AppDaos
	config     interfaces.AppConfig
	logger     *zap.Logger
	services   interfaces.AppServices
	router     *mux.Router
	rootLogger *zap.Logger
}

func (a AppImpl) Config() interfaces.AppConfig {
	return a.config
}

func (a AppImpl) Dao() interfaces.AppDaos {
	return a.daos
}

func (a AppImpl) DB() *gorm.DB {
	return a.db
}

func (a AppImpl) Logger() *zap.Logger {
	return a.logger
}

func (a AppImpl) Service() interfaces.AppServices {
	return a.services
}

func (a AppImpl) Router() *mux.Router {
	return a.router
}

func (a AppImpl) RootLogger() *zap.Logger {
	return a.rootLogger
}

func NewApp(cfg interfaces.AppConfig, db *gorm.DB, daos interfaces.AppDaos, svcs interfaces.AppServices,
	router *mux.Router, appLogger types.AppLogger, rootLogger types.RootLogger) *AppImpl {
	return &AppImpl{
		db:         db,
		daos:       daos,
		config:     cfg,
		logger:     appLogger,
		services:   svcs,
		router:     router,
		rootLogger: rootLogger,
	}
}

// AppServicesImpl Implementation of the interfaces.AppServices interface for exposing services to the application.
type AppServicesImpl struct {
	jwt   interfaces.JWTService
	login interfaces.LoginService
}

func (a AppServicesImpl) Login() interfaces.LoginService {
	return a.login
}

func (a AppServicesImpl) JWT() interfaces.JWTService {
	return a.jwt
}

func NewAppServices(jwt interfaces.JWTService, login interfaces.LoginService) *AppServicesImpl {
	return &AppServicesImpl{
		jwt:   jwt,
		login: login,
	}
}

// AppDaosImpl Implementation of the interfaces.AppDaos interface for exposing DAOs to the application.
type AppDaosImpl struct {
	users interfaces.UserDao
}

func (a AppDaosImpl) Users() interfaces.UserDao {
	return a.users
}

func NewAppDaos(users interfaces.UserDao) *AppDaosImpl {
	return &AppDaosImpl{
		users: users,
	}
}

// AppConfigImpl Implementation of the interfaces.AppConfig interface for application configuration.
type AppConfigImpl struct {
	autoMigrate bool
	env         string
	listen      interfaces.ListenConfig
	mysql       interfaces.MySQLConfig
}

func (a AppConfigImpl) AutoMigrate() bool {
	return a.autoMigrate
}

func (a AppConfigImpl) Env() string {
	return a.env
}

func (a AppConfigImpl) Listen() interfaces.ListenConfig {
	return a.listen
}

func (a AppConfigImpl) MySQL() interfaces.MySQLConfig {
	return a.mysql
}

type AppConfigAutoMigrate bool
type AppConfigEnv string

func NewAppConfig(autoMigrate AppConfigAutoMigrate, env AppConfigEnv, listen interfaces.ListenConfig,
	mysql interfaces.MySQLConfig) *AppConfigImpl {
	return &AppConfigImpl{
		autoMigrate: bool(autoMigrate),
		env:         string(env),
		listen:      listen,
		mysql:       mysql,
	}
}
