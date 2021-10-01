// +build wireinject

package pkg

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/dao"
	"github.com/naftulikay/golang-webapp/pkg/database"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/service"
	"github.com/naftulikay/golang-webapp/pkg/types"
)

func initializeApp(
	appCfg interfaces.AppConfig,
	jwtKey types.JWTKey,
	rootLogger types.RootLogger,
) (interfaces.App, error) {
	wire.Build(
		// interfaces
		wire.Bind(new(interfaces.App), new(*service.AppImpl)),
		wire.Bind(new(interfaces.AppServices), new(*service.AppServicesImpl)),
		wire.Bind(new(interfaces.AppDaos), new(*service.AppDaosImpl)),
		wire.Bind(new(interfaces.JWTService), new(*service.JWTServiceImpl)),
		wire.Bind(new(interfaces.LoginService), new(*service.LoginServiceImpl)),
		wire.Bind(new(interfaces.UserDao), new(*dao.UserDaoImpl)),
		// providers
		service.NewApp,
		service.NewAppServices,
		service.NewAppDaos,
		// providers/services
		service.NewJWTService,
		service.NewLoginService,
		// providers/daos
		dao.NewUserDao,
		// providers/logging
		types.NewAppLogger,
		types.NewJWTServiceLogger,
		types.NewLoginServiceLogger,
		types.NewUserDaoLogger,
		types.NewDatabaseLogger,
		// providers/database
		database.Connect,
		// providers/external
		mux.NewRouter,
		// extractors
		database.ExtractMySQLConfig,
	)

	return nil, nil
}
