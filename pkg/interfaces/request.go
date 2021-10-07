package interfaces

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"go.uber.org/zap"
	"net/http"
)

type RequestContext interface {
	// ID The unique request ID as an UUID.
	ID() uuid.UUID
	// App The top-level application instance, which provides access to application services and configuration.
	App() App
	// Config Read application-level configuration.
	Config() AppConfig
	// Services Access the different services available to the application.
	Services() AppServices
	// Dao Access the different data-access-objects available to the application.
	Dao() AppDaos
	// Logger The unique, per-handler logger for the request.
	Logger() *zap.Logger
	// Req The raw http.Request object.
	Req() *http.Request
	// Context The request context.
	Context() context.Context
	// Authenticated Whether the current request has valid auth data.
	Authenticated() bool
	// IsUser Whether the current request is authorized for user access.
	IsUser() bool
	// IsAdmin Whether the current request is authorized for admin access.
	IsAdmin() bool
	// Token Return the JWT token if the current request has a valid one.
	Token() *jwt.Token
	// Claims Return the JWT claims if the current request has valid JWT.
	Claims() *auth.JWTClaims
	// JSON Unmarshal the JSON request body.
	JSON(dest interface{}) error
	// JSONV Unmarshal the JSON request body and run validator.Validate on it.
	JSONV(dest interface{}) error

	// FIXME implement tracing
}
