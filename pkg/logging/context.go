package logging

import (
	"context"
	"github.com/google/uuid"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"go.uber.org/zap"
)

// ZapFieldsFromContext Extract public fields from the context and return them as a list of zap.Field objects.
//
// This is to make structured logging work well. In the body of request handlers, the following code can be used to
// attach request-related data to the route logger:
//
// ```
// func LoginHandler(req interfaces.RequestContext) *views.Response {
//     logger := req.Logger().With(logging.ZapFieldsFromContext(req.Context()))
//     // now, the request id and other metadata will automatically be included in future logs
// }
// ```
func ZapFieldsFromContext(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0)

	// attach request id
	if ctx.Value(constants.ContextKeyRequestID) != nil {
		id, ok := ctx.Value(constants.ContextKeyRequestID).(uuid.UUID)

		if !ok {
			panic("unable to get request id as uuid from context")
		}

		fields = append(fields, zap.String(constants.ContextKeyRequestID, id.String()))
	} else {
		panic("request id not present in the context")
	}

	// attach authenticated
	if ctx.Value(constants.ContextKeyAuthenticated) != nil {
		authenticated, ok := ctx.Value(constants.ContextKeyAuthenticated).(bool)

		if !ok {
			panic("authenticated field not present in the context")
		} else {
			fields = append(fields, zap.Bool("authenticated", authenticated))
		}

		if authenticated {
			if ctx.Value(constants.ContextKeyUserID) != nil {
				userID, ok := ctx.Value(constants.ContextKeyUserID).(uint64)

				if !ok {
					panic("user_id not present in the context")
				}

				fields = append(fields, zap.Uint64("user_id", userID))
			}

			if ctx.Value(constants.ContextKeyUserEmail) != nil {
				userEmail, ok := ctx.Value(constants.ContextKeyUserEmail).(string)

				if !ok {
					panic("user_email not present in the context")
				}

				fields = append(fields, zap.String("user_email", userEmail))
			}

			if ctx.Value(constants.ContextKeyUserRole) != nil {
				userRole, ok := ctx.Value(constants.ContextKeyUserRole).(string)

				if !ok {
					panic("user_role not present in the context")
				}

				fields = append(fields, zap.String("user_role", userRole))
			}
		}
	}

	return fields
}
