package handlers

import (
	"github.com/gorilla/handlers"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// CORSAllowedHeaders Return a list of allowed CORS headers.
func CORSAllowedHeaders() []string {
	return []string{
		// Accept is implicit
		// Accept-Language is implicit
		// Content-Language is implicit
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Connection",
		"Pragma",
		"Referer",
		"Sec-Fetch-Dest",
		"Sec-Fetch-Mode",
		"Sec-Fetch-Site",
		"User-Agent",
		"X-Requested-With",
	}
}

// CORSAllowedMethods Return a list of allowed CORS methods.
func CORSAllowedMethods() []string {
	return []string{
		"GET", "HEAD", "POST", "PUT", "PATCH", "PUT", "OPTIONS",
	}
}

// CORSOriginValidator Generates a CORS origin validator function.
func CORSOriginValidator(config interfaces.AppConfig, logger *zap.Logger) func(string) bool {
	return func(origin string) bool {
		switch env := config.Env(); env {
		case constants.DevEnvironment:
			value, err := url.Parse(origin)

			if err != nil {
				logger.Warn("Unable to parse origin during CORS origin validation, denying request.",
					zap.Error(err), zap.String("origin", origin))

				return false
			}

			// if the hostname is localhost, or it ends in .local, or it is a loopback IP address, allow
			isLocalhost := value.Hostname() == "localhost"
			isAtLocalDomain := strings.HasSuffix(value.Hostname(), ".local")
			isLoopbackIP := false

			if !isLocalhost && !isAtLocalDomain {
				// now we attempt to parse an IP in case the hostname is an IP address
				ip := net.ParseIP(value.Hostname())

				if ip != nil {
					isLoopbackIP = ip.IsLoopback()
				}
			}

			if isLocalhost || isAtLocalDomain || isLoopbackIP {
				logger.Debug("Allowing request from origin.", zap.String("origin", origin))

				return true
			} else {
				logger.Warn("Denying request from unknown origin.", zap.String("origin", origin),
					zap.String("hostname", value.Hostname()))

				return false
			}
		case constants.StagingEnvironment:
			logger.Error("CORS origin validation is not implemented for the staging environment, denying request.")
			return false
		case constants.ProductionEnvironment:
			logger.Error("CORS origin validation is not implemented for the production environment, denying request.")
			return false
		default:
			logger.Error("Unknown environment, denying request.", zap.String("env", config.Env()))
			return false
		}
	}
}

// CORS Returns a http.Handler wrapper function for adding CORS validation to an application.
func CORS(config interfaces.AppConfig, logger *zap.Logger) func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedHeaders(CORSAllowedHeaders()),
		handlers.AllowedMethods(CORSAllowedMethods()),
		handlers.AllowedOriginValidator(CORSOriginValidator(config, logger)),
	)
}
