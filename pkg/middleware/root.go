package middleware

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
)

func ConfigureRoot(app interfaces.App) {
	app.Router().Use(RequestMetadata())
}
