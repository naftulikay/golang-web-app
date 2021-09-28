package pkg

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/database"
	"github.com/naftulikay/golang-webapp/pkg/handlers"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/logging"
	"github.com/naftulikay/golang-webapp/pkg/middleware"
	"github.com/naftulikay/golang-webapp/pkg/routes"
	_ "github.com/swaggo/swag"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

func Start(config interfaces.AppConfig) {
	// initialize logging
	rootLogger, err := logging.Configure(config)
	rootLogger = rootLogger.Named("app")

	if err != nil {
		log.Fatalf("Unable to configure Zap logging: %s", err)
	}

	// connect to database
	db, err := database.Connect(config.MySQL(), rootLogger.Named("database"))

	if err != nil {
		rootLogger.Fatal("Unable to connect to database.")
	}

	// migrate models
	if config.AutoMigrate() {
		err = database.AutoMigrate(db, rootLogger.Named("migrator"))

		if err != nil {
			rootLogger.Warn("Automatic database migration failed, proceeding anyway.",
				zap.Error(err))
		}
	}

	// FIXME wire together app

	// setup router
	router := mux.NewRouter()

	// setup middleware
	middleware.ConfigureRoot(nil)

	// configure routes
	routes.ConfigureRoutes(nil, router)

	// setup http server
	listenAddr := fmt.Sprintf("%s:%d", config.Listen().Host(), config.Listen().Port())

	rootLogger.Info("Starting web server",
		zap.String("url", fmt.Sprintf("http://%s", listenAddr)),
		zap.String("host", config.Listen().Host()),
		zap.Uint16("port", config.Listen().Port()))

	server := &http.Server{
		Addr: listenAddr,
		// wrap the entire router in CORS because CORS needs access to everything basically
		Handler:      handlers.CORS(config, rootLogger.Named("cors"))(router),
		ReadTimeout:  300 * time.Second, // support long debugging sessions
		WriteTimeout: 300 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to serve application: %s", err)
	}
}
