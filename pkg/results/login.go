package results

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
)

func NewLoginResult(user models.User, signedToken string, token jwt.Token, claims auth.JWTClaims) interfaces.LoginResult {
	return LoginResultImpl{
		user:        user,
		signedToken: signedToken,
		token:       token,
		claims:      claims,
	}
}

type LoginResultImpl struct {
	user        models.User
	signedToken string
	token       jwt.Token
	claims      auth.JWTClaims
}

func (l LoginResultImpl) User() models.User {
	return l.user
}

func (l LoginResultImpl) SignedToken() string {
	return l.signedToken
}

func (l LoginResultImpl) Token() jwt.Token {
	return l.token
}

func (l LoginResultImpl) Claims() auth.JWTClaims {
	return l.claims
}
