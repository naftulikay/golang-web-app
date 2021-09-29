package interfaces

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/models"
)

type LoginResult interface {
	User() *models.User
	SignedToken() *string
	Token() *jwt.Token
	Claims() *auth.JWTClaims
}
