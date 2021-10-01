package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/request"
	"github.com/naftulikay/golang-webapp/pkg/views"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

var routePrefix = regexp.MustCompile(`^[/]?pkg[/\.]routes[/\.]`)

// APIHandlerFunc A HTTP handler which conforms to our API framework.
type APIHandlerFunc func(req interfaces.RequestContext) *views.Response

func getRouteLoggerName(handler APIHandlerFunc) string {
	// get the fully qualified function name (github.com/naftulikay/golang-webapp/pkg/routes/admin/users.CreateUserHandler)
	fq := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	// strip out the module name (/pkg/routes/admin/users.CreateUserHandler)
	fq = strings.Replace(fq, constants.GoModule, "", 1)
	// strip the prefix (admin/users.CreateUserHandler)
	fq = routePrefix.ReplaceAllString(fq, "")
	// replace all slashes with dots (admin.users.CreateUserHandler)
	fq = strings.ReplaceAll(fq, "/", ".")

	return fq
}

func API(app interfaces.App, handler APIHandlerFunc) http.HandlerFunc {
	return apiWithLoggerName(app, getRouteLoggerName(handler), handler)
}

func apiWithLoggerName(app interfaces.App, loggerName string, handler APIHandlerFunc) http.HandlerFunc {
	logger := app.Logger().Named(fmt.Sprintf("routes.%s", loggerName))

	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME add tracing span
		req, err := request.NewRequestContext(app, logger, r)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("Unable to build request context for request.", zap.Error(err))
			return
		}

		// set content type
		w.Header().Set("Content-Type", "application/json")

		// execute the handler
		resp := handler(req)

		if resp == nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("API handler returned a nil result.")
			return
		}

		// set the response code from the response
		w.WriteHeader(int(resp.StatusCode))

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Warn("Unable to write response body, did client terminate connection?", zap.Error(err))
		}
	}
}
