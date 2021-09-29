package interfaces

import (
	"github.com/naftulikay/golang-webapp/pkg/models"
)

type JWTService interface {
	// Generate Create and sign a token for a given user, returning the signed token string, the claims, and optionally
	// an error if something went wrong.
	Generate(user *models.User) (*JWTGenerateResult, error)
	// Validate Securely validate a token string, returning the parsed token and its claims.
	//
	// If this function returns the token and claims (e.g. if both are not nil), the API contract must be that
	// token.Valid is true. If either the token or the claims are nil, or the error is not nil, validation has failed
	// either in an expected or unexpected way.
	Validate(token string) (*JWTValidateResult, error)
}

type LoginService interface {
	Login(email, password string) (*LoginResult, error)
}
