package results

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
)

func NewJWTGenerateResult(signedToken *string, token *jwt.Token, claims *auth.JWTClaims) interfaces.JWTGenerateResult {
	return JWTGenerateResultImpl{
		signedToken: *signedToken,
		token:       *token,
		claims:      *claims,
	}
}

type JWTGenerateResultImpl struct {
	signedToken string
	token       jwt.Token
	claims      auth.JWTClaims
}

func (j JWTGenerateResultImpl) SignedToken() string {
	return j.signedToken
}

func (j JWTGenerateResultImpl) Token() jwt.Token {
	return j.token
}

func (j JWTGenerateResultImpl) Claims() auth.JWTClaims {
	return j.claims
}

func NewJWTValidateResult(token jwt.Token, claims auth.JWTClaims) interfaces.JWTValidateResult {
	return JWTValidateResultImpl{
		token:  token,
		claims: claims,
	}
}

type JWTValidateResultImpl struct {
	token  jwt.Token
	claims auth.JWTClaims
}

func (j JWTValidateResultImpl) Token() jwt.Token {
	return j.token
}

func (j JWTValidateResultImpl) Claims() auth.JWTClaims {
	return j.claims
}
