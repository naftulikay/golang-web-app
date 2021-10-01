package pkg

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/database"
	"github.com/naftulikay/golang-webapp/pkg/handlers"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/logging"
	"github.com/naftulikay/golang-webapp/pkg/middleware"
	"github.com/naftulikay/golang-webapp/pkg/routes"
	"github.com/naftulikay/golang-webapp/pkg/types"
	_ "github.com/swaggo/swag"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

func Start(config interfaces.AppConfig, jwtKey types.JWTKey) {
	// initialize logging
	rootLogger, err := logging.Configure(config)
	logger := rootLogger.Named("app")

	if err != nil {
		log.Fatalf("Unable to configure Zap logging: %s", err)
	}

	app, err := initializeApp(
		config,
		jwtKey,
		rootLogger,
	)

	// migrate models
	if config.AutoMigrate() {
		err = database.AutoMigrate(app.DB(), logger.Named("migrator"))

		if err != nil {
			logger.Warn("Automatic database migration failed, proceeding anyway.",
				zap.Error(err))
		}
	}

	// setup middleware
	middleware.ConfigureRoot(app)

	// configure routes
	routes.ConfigureRoutes(app, app.Router())

	// setup http server
	listenAddr := fmt.Sprintf("%s:%d", config.Listen().Host(), config.Listen().Port())

	rootLogger.Info("Starting web server",
		zap.String("url", fmt.Sprintf("http://%s", listenAddr)),
		zap.String("host", config.Listen().Host()),
		zap.Uint16("port", config.Listen().Port()))

	server := &http.Server{
		Addr: listenAddr,
		// wrap the entire router in CORS because CORS needs access to everything basically
		Handler:      handlers.CORS(config, logger.Named("cors"))(app.Router()),
		ReadTimeout:  300 * time.Second, // support long debugging sessions
		WriteTimeout: 300 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to serve application: %s", err)
	}
}
