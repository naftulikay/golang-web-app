package constants

const (
	// ContextKeyRequestID Context key for the uuid.UUID request ID attached to each request.
	ContextKeyRequestID = "request_id"
	// ContextKeyAuthenticated Context key for the boolean `authenticated` context value.
	ContextKeyAuthenticated = "authenticated"
	// ContextKeyJWTToken Context key for the jwt.Token associated with a request, if any.
	ContextKeyJWTToken = "jwt_token"
	// ContextKeyJWTClaims Context key for the auth.JWTClaims object associated with a request, if any.
	ContextKeyJWTClaims = "jwt_claims"
	// ContextKeyUserID Context key for the uint user ID associated with a request, if any.
	ContextKeyUserID = "user_id"
	// ContextKeyUserEmail Context key for the string user email associated with a request, if any.
	ContextKeyUserEmail = "user_email"
	// ContextKeyUserRole Context key for the string user's role associated with a request, if any.
	ContextKeyUserRole = "user_role"
)
