package middleware

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
)

func ConfigureRoot(app interfaces.App) {
	app.Router().Use(RequestMetadata())
}

func ConfigureApi(app interfaces.App, router *mux.Router) {

}
