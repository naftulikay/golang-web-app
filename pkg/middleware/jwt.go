package middleware

import (
	"context"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"
)

var bearer = regexp.MustCompile(`^Bearer\s+`)

// JWTAuth A middleware function for extracting, parsing, and validating JWT auth data from request headers.
//
// This middleware sets `authenticated` to either `true` or `false`, depending on whether the token exists, is valid,
// and is not expired. The next most interesting field is `user_role`, which dictates the permission level of the
// current user. JWT tokens and claims are attached as well, in addition to the user's ID and email.
func JWTAuth(app interfaces.App) func(http.Handler) http.Handler {
	logger := app.Logger().Named("middleware.jwt")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractTokenFromRequest(r)

			ctx := r.Context()

			if token != nil {
				logger.Debug("Token extraction succeeded, attempting to validate token.")
				result, err := app.Service().JWT().Validate(*token)

				if err != nil {
					// jwt validation failed, so set authenticated to false
					logger.Debug("JWT token validation failed, marking request as unauthenticated.", zap.Error(err))

					ctx = context.WithValue(ctx, constants.ContextKeyAuthenticated, false)
				} else {
					// jwt validation succeeded, attach metadata
					logger.Debug("JWT token validation succeeded, marking request as authenticated and adding metadata.")

					ctx = context.WithValue(ctx, constants.ContextKeyAuthenticated, true)
					ctx = context.WithValue(ctx, constants.ContextKeyJWTToken, result.Token())
					ctx = context.WithValue(ctx, constants.ContextKeyJWTClaims, result.Claims())
					ctx = context.WithValue(ctx, constants.ContextKeyUserID, result.Claims().UserID)
					ctx = context.WithValue(ctx, constants.ContextKeyUserEmail, result.Claims().Email)
					ctx = context.WithValue(ctx, constants.ContextKeyUserRole, result.Claims().Role)
				}
			} else {
				// unable to extract jwt token from Authorization header, set authenticated to false
				logger.Debug("JWT token data missing from Authorization header, marking request as unauthenticated.")

				ctx = context.WithValue(ctx, constants.ContextKeyAuthenticated, false)
			}

			// update the request with the updated context
			*r = *r.WithContext(ctx)

			// kick the next metadata service
			next.ServeHTTP(w, r)
		})
	}
}

func extractTokenFromRequest(r *http.Request) *string {
	header := r.Header.Get("Authorization")

	if len(header) == 0 {
		// header doesn't exist or is empty
		return nil
	}

	// check if we even have a bearer token:
	if !strings.HasPrefix(header, "Bearer ") {
		return nil
	}

	token := strings.TrimSpace(bearer.ReplaceAllString(header, ""))

	if len(token) == 0 {
		return nil
	}

	return &token
}
