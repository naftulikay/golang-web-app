package interfaces

import (
	"context"
	"github.com/google/uuid"
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
	// JSON Unmarshal the JSON request body.
	JSON(dest interface{}) error
	// JSONV Unmarshal the JSON request body and run validator.Validate on it.
	JSONV(dest interface{}) error

	// FIXME implement tracing
}
