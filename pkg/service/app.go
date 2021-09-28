package service

import "github.com/naftulikay/golang-webapp/pkg/interfaces"

type AppImpl struct {
	config   interfaces.AppConfig
	services interfaces.AppServices
	daos     interfaces.AppDaos
}

func (a AppImpl) Config() interfaces.AppConfig {
	return a.config
}

func (a AppImpl) Dao() interfaces.AppDaos {
	return a.daos
}

func (a AppImpl) Service() interfaces.AppServices {
	return a.services
}

func NewApp(params AppParams) AppImpl {
	return AppImpl{}
}

type AppServicesImpl struct {
	login interfaces.LoginService
}

func (a AppServicesImpl) Login() interfaces.LoginService {
	return a.login
}

func NewAppServices(params AppServicesParams) AppServicesImpl {
	return AppServicesImpl{}
}

type AppDaosImpl struct{}

func NewAppDaos(params AppDaosParams) AppDaosImpl {
	return AppDaosImpl{}
}
