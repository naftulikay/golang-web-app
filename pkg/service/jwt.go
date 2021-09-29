package service

import (
	"crypto/subtle"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"github.com/naftulikay/golang-webapp/pkg/utils"
	"time"
)

const (
	JWTIssuer = "github.com/nafutlikay/golang-webapp"
	JWTExpiry = 30 * 24 * time.Hour
)

type JWTServiceImpl struct {
	key [32]byte
}

func (j JWTServiceImpl) safe() bool {
	zero := utils.NullBytes32()

	// j.key MUST NOT be equal to a completely zeroed array
	return subtle.ConstantTimeCompare(j.key[:], zero[:]) != 1
}

func (j JWTServiceImpl) secretFactory(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return j.key[:], nil
}

func (j JWTServiceImpl) Generate(user *models.User) (*interfaces.JWTGenerateResult, error) {
	if !j.safe() {
		panic("jwt key has a zero value")
	}

	now := time.Now().UTC()

	claims := auth.JWTClaims{
		UserID:    uint64(user.ID),
		Email:     user.Email,
		Role:      user.Role,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWTIssuer,
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(JWTExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(j.key[:])

	if err != nil {
		return nil, err
	}

	result := newJWTGenerateResult(&signed, token, &claims)

	return &result, nil
}

func (j JWTServiceImpl) Validate(encodedToken string) (*interfaces.JWTValidateResult, error) {
	if !j.safe() {
		panic("jwt key has a zero value")
	}

	token, err := jwt.ParseWithClaims(encodedToken, &auth.JWTClaims{}, j.secretFactory)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*auth.JWTClaims); ok && token.Valid {
		result := newJWTValidateResult(token, claims)
		return &result, nil
	}

	// this occurs when we've got claims of the wrong type or when the token is invalid
	return nil, nil
}

// jwt generate result implementation
func newJWTGenerateResult(signedToken *string, token *jwt.Token, claims *auth.JWTClaims) interfaces.JWTGenerateResult {
	return jwtGenerateResult{
		signedToken: signedToken,
		token:       token,
		claims:      claims,
	}
}

type jwtGenerateResult struct {
	signedToken *string
	token       *jwt.Token
	claims      *auth.JWTClaims
}

func (j jwtGenerateResult) SignedToken() *string {
	return j.signedToken
}

func (j jwtGenerateResult) Token() *jwt.Token {
	return j.token
}

func (j jwtGenerateResult) Claims() *auth.JWTClaims {
	return j.claims
}

// jwt validate result implementation
func newJWTValidateResult(token *jwt.Token, claims *auth.JWTClaims) interfaces.JWTValidateResult {
	return jwtValidateResult{
		token:  token,
		claims: claims,
	}
}

type jwtValidateResult struct {
	token  *jwt.Token
	claims *auth.JWTClaims
}

func (j jwtValidateResult) Token() *jwt.Token {
	return j.token
}

func (j jwtValidateResult) Claims() *auth.JWTClaims {
	return j.claims
}
