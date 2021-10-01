package request

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
)

type JWTData struct {
	SignedToken string
	Token       jwt.Token
	Claims      auth.JWTClaims
}
