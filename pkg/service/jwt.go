package service

import (
	"crypto/subtle"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"github.com/naftulikay/golang-webapp/pkg/results"
	"github.com/naftulikay/golang-webapp/pkg/utils"
	"go.uber.org/zap"
	"time"
)

const (
	JWTIssuer = "github.com/nafutliikay/golang-webapp"
	JWTExpiry = 30 * 24 * time.Hour
)

type JWTServiceImpl struct {
	key    [32]byte
	logger *zap.Logger
}

func NewJWTService(key [32]byte, logger *zap.Logger) (*JWTServiceImpl, error) {
	r := JWTServiceImpl{
		key:    key,
		logger: logger,
	}

	if !r.safe() {
		return nil, fmt.Errorf("key must not be a null array")
	}

	return &r, nil
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

	result := results.NewJWTGenerateResult(&signed, token, &claims)

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

	claims, claimsOk := token.Claims.(*auth.JWTClaims)

	if claimsOk && token.Valid {
		j.logger.Debug("Token is valid and claims are of the correct type.",
			zap.Uint64("user_id", claims.UserID), zap.String("user_email", claims.Email))

		result := results.NewJWTValidateResult(*token, *claims)
		return &result, nil
	}

	if !claimsOk {
		j.logger.Debug("Received invalid claims object, unable to deserialize as auth.JWTClaims.")

		return nil, fmt.Errorf("invalid claims")
	}

	// this occurs when we've got claims of the wrong type or when the token is invalid
	return nil, nil
}
