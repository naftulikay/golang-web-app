// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/naftulikay/golang-webapp/pkg/dao"
	"github.com/naftulikay/golang-webapp/pkg/database"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/service"
	"github.com/naftulikay/golang-webapp/pkg/types"
)

func initializeLoginService(config interfaces.MySQLConfig, jwtKey types.JWTKey,
	loginServiceLogger types.LoginServiceLogger, jwtServiceLogger types.JWTServiceLogger,
	userDaoLogger types.UserDaoLogger, databaseLogger types.DatabaseLogger) (interfaces.LoginService, error) {

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
