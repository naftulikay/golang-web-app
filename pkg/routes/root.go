package routes

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/middleware"
	"github.com/naftulikay/golang-webapp/pkg/routes/test"

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

	// attach JWT auth parsing/validation middleware to anything within /api
	api.Use(middleware.JWTAuth(app))

	v1 := api.PathPrefix("/v1").Subrouter()

	login.ConfigureRoutes(app, v1)
	test.ConfigureRoutes(app, v1.PathPrefix("/test").Subrouter())
}
