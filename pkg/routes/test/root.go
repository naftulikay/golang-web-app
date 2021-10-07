package test

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/handlers"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/response"
	"github.com/naftulikay/golang-webapp/pkg/views"
	"net/http"
)

func ConfigureRoutes(app interfaces.App, r *mux.Router) {
	r.Path("/is-user").HandlerFunc(handlers.UserAPI(app, UserHandler)).Methods("GET")
	r.Path("/is-admin").HandlerFunc(handlers.AdminAPI(app, AdminHandler)).Methods("GET")
}

// UserHandler godoc
// @Summary Test user privileges.
// @Tags test
// @Description Tests whether a user is logged-in.
// @Produce json
// @Success 200 {object} views.Response
// @Failure 403 {object} views.Response
// @Security JWT
// @Router /test/is-user [get]
func UserHandler(req interfaces.RequestContext) *views.Response {
	req.Logger().Debug("User logged in!")
	return response.Ok(req, nil)
}

// AdminHandler godoc
// @Summary Test admin privileges.
// @Tags test
// @Description Tests whether a user is logged-in and an admin.
// @Produce json
// @Success 200 {object} views.Response
// @Failure 403 {object} views.Response
// @Security JWT
// @Router /test/is-admin [get]
func AdminHandler(req interfaces.RequestContext) *views.Response {
	req.Logger().Debug("Admin logged in!")

	return &views.Response{
		StatusCode: http.StatusOK,
		Links: views.ResponseLinks{
			Self: req.Req().URL.Path,
		},
	}
}
