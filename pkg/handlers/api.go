package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/request"
	"github.com/naftulikay/golang-webapp/pkg/response"
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

type AuthorizerFunc func(req interfaces.RequestContext) bool

func UserAuthorizer() AuthorizerFunc {
	return func(req interfaces.RequestContext) bool {
		return req.Authenticated() && req.IsUser()
	}
}

func AdminAuthorizer() AuthorizerFunc {
	return func(req interfaces.RequestContext) bool {
		return req.Authenticated() && req.IsAdmin()
	}
}

func AnonAuthorizer() AuthorizerFunc {
	return func(req interfaces.RequestContext) bool {
		return true
	}
}

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

// API Create a handler which is accessible by any user, authenticated or not.
func API(app interfaces.App, handler APIHandlerFunc) http.HandlerFunc {
	return buildAPIHandler(app, getRouteLoggerName(handler), handler, AnonAuthorizer())
}

// UserAPI Create a handler which is accessible only by authenticated users and admins.
func UserAPI(app interfaces.App, handler APIHandlerFunc) http.HandlerFunc {
	return buildAPIHandler(app, getRouteLoggerName(handler), handler, UserAuthorizer())
}

// AdminAPI Create a handler which is accessible only by authenticated admins.
func AdminAPI(app interfaces.App, handler APIHandlerFunc) http.HandlerFunc {
	return buildAPIHandler(app, getRouteLoggerName(handler), handler, AdminAuthorizer())
}

func buildAPIHandler(app interfaces.App, loggerName string, handler APIHandlerFunc, authorizer AuthorizerFunc) http.HandlerFunc {
	logger := app.Logger().Named(fmt.Sprintf("routes.%s", loggerName))

	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME add tracing span
		req, err := request.NewRequestContext(app, logger, r)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("Unable to build request context for request.", zap.Error(err))
			return
		}

		// authenticate/authorize before execution and terminate before reaching the handler
		if !authorizer(req) {
			req.Logger().Debug("Access forbidden.")
			w.WriteHeader(http.StatusForbidden)

			if err := json.NewEncoder(w).Encode(response.Forbidden(req, "Access denied.")); err != nil {
				logger.Warn("Unable to write response body, did client terminate connection?", zap.Error(err))
			}

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
