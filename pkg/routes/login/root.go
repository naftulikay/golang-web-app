package login

import (
	"github.com/gorilla/mux"
	"github.com/naftulikay/golang-webapp/pkg/handlers"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/response"
	"github.com/naftulikay/golang-webapp/pkg/views"
	"go.uber.org/zap"
	"time"
)

func ConfigureRoutes(app interfaces.App, r *mux.Router) {
	r.HandleFunc("/login", handlers.API(app, LoginHandler)).Methods("POST")
}

// LoginHandler godoc
// @Summary Login
// @Tags auth
// @Description Log a user in using email and password.
// @Accept json
// @Param login body views.LoginRequest true "Login Credentials"
// @Produce json
// @Success 200 {object} views.Response{data=views.LoginResponse}
// @Failure 400 {object} views.Response
// @Failure 403 {object} views.Response
// @Router /login [post]
func LoginHandler(req interfaces.RequestContext) *views.Response {
	var reqBody views.LoginRequest

	if err := req.JSONV(&reqBody); err != nil {
		return response.BadRequest(req, err.Error())
	}

	// okay we've got a good request, so do the login
	result, err := req.Services().Login().Login(reqBody.Email, reqBody.Password)

	if err != nil {
		req.Logger().Debug("Login failed.", zap.String("user_email", reqBody.Email))
		return response.Forbidden(req, "Login failed.")
	}

	req.Logger().Debug("Login succeeded.", zap.Uint("user_id", (*result).User().ID),
		zap.String("user_email", (*result).User().Email))

	return response.Ok(req, views.LoginResponse{
		UserID:    (*result).User().ID,
		Email:     (*result).User().Email,
		FirstName: (*result).User().FirstName,
		LastName:  (*result).User().LastName,
		Role:      (*result).User().Role,
		Token:     (*result).SignedToken(),
		IssuedAt:  time.Unix((*result).Claims().IssuedAt, 0).UTC(),
		NotBefore: time.Unix((*result).Claims().NotBefore, 0).UTC(),
		ExpiresAt: time.Unix((*result).Claims().ExpiresAt, 0).UTC(),
	})
}
