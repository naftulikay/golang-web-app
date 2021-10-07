package interfaces

import (
	"github.com/naftulikay/golang-webapp/pkg/models"
)

type JWTService interface {
	// Generate Create and sign a token for a given user.
	Generate(user *models.User) (JWTGenerateResult, error)
	// Validate Securely validate a token string, returning the relevant parsed JWT data.
	//
	// If JWTValidateResult is nil, validation has failed for the token. If err is not nil,
	Validate(token string) (JWTValidateResult, error)
}

type LoginService interface {
	Login(email, password string) (*LoginResult, error)
}
