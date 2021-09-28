package pkg

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"log"
)

func Start(config interfaces.AppConfig) {
	log.Printf("Environment: %s", config.Env())

	log.Printf("Dumping listen configuration...")
	log.Printf("listen_host: %s, listen_port: %d", config.Listen().Host(), config.Listen().Port())

	log.Printf("Dumping MySQL configuration...")
	log.Printf("mysql_host: %s, mysql_port: %d, mysql_database: %s, mysql_user: %s, mysql_password: %s",
		config.MySQL().Host(), config.MySQL().Port(), config.MySQL().Database(), config.MySQL().User(),
		config.MySQL().Password())
}
