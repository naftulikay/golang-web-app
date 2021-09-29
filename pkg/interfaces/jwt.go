package interfaces

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
)

type JWTGenerateResult interface {
	SignedToken() string
	Token() jwt.Token
	Claims() auth.JWTClaims
}

type JWTValidateResult interface {
	Token() jwt.Token
	Claims() auth.JWTClaims
}
