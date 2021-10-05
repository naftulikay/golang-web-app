package routes

import (
	"github.com/gorilla/mux"
	// generated code from `swag init`, ignore it during security auditing
	_ "github.com/naftulikay/golang-webapp/docs" // #nosec
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/routes/login"
	swagger "github.com/swaggo/http-swagger"
)

func ConfigureRoutes(app interfaces.App, router *mux.Router) {
	// bind swagger endpoints
	router.PathPrefix("/swagger/").Handler(swagger.WrapHandler)

	api := router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	login.ConfigureRoutes(app, v1)
}
