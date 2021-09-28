package pkg

import (
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Start(config interfaces.AppConfig) {
	log.Printf("Environment: %s", config.Env())

	log.Printf("Dumping listen configuration...")
	log.Printf("listen_host: %s, listen_port: %d", config.Listen().Host(), config.Listen().Port())

	log.Printf("Dumping MySQL configuration...")
	log.Printf("mysql_host: %s, mysql_port: %d, mysql_database: %s, mysql_user: %s, mysql_password: %s",
		config.MySQL().Host(), config.MySQL().Port(), config.MySQL().Database(), config.MySQL().User(),
		config.MySQL().Password())

	var db *gorm.DB

	// connect to database
	err := backoff.Retry(func() error {
		uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", config.MySQL().User(),
			config.MySQL().Password(), config.MySQL().Host(), config.MySQL().Port(), config.MySQL().Database())

		connection, err := gorm.Open(mysql.Open(uri))

		if err != nil {
			return err
		}

		db = connection

		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 15))

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	log.Printf("Database connected (%v)", db)
	log.Printf("Initialization complete, serving at http://%s:%d/", config.Listen().Host(),
		config.Listen().Port())

	// migrate models
	err = db.AutoMigrate(models.User{})

	if err != nil {
		log.Fatalf("Unable to migrate models: %s", err)
	}

	// setup http server
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/get", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request!")
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Listen().Host(), config.Listen().Port()),
		Handler:      mux,
		ReadTimeout:  300 * time.Second, // support long debugging sessions
		WriteTimeout: 300 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to serve application: %s", err)
	}
}
