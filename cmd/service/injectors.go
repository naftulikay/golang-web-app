// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/naftulikay/golang-webapp/pkg/dao"
	"github.com/naftulikay/golang-webapp/pkg/database"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/service"
)

func initializeLoginService(config interfaces.MySQLConfig, jwtKey service.JWTKey,
	loginServiceLogger service.LoginServiceLogger, jwtServiceLogger service.JWTServiceLogger,
	userDaoLogger dao.UserDaoLogger, databaseLogger database.DatabaseLogger) (interfaces.LoginService, error) {

	wire.Build(
		// interfaces
		wire.Bind(new(interfaces.LoginService), new(*service.LoginServiceImpl)),
		wire.Bind(new(interfaces.JWTService), new(*service.JWTServiceImpl)),
		wire.Bind(new(interfaces.UserDao), new(*dao.UserDaoImpl)),
		// services
		service.NewLoginService,
		service.NewJWTService,
		// daos
		dao.NewUserDao,
		// database
		database.Connect,
	)

	return nil, nil
}
