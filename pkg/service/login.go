package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/auth"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"go.uber.org/zap"
)

type LoginServiceImpl struct {
	dao    interfaces.UserDao
	jwt    interfaces.JWTService
	logger *zap.Logger
}

func (l LoginServiceImpl) Login(email, password string) (*interfaces.LoginResult, error) {
	loginFailure := fmt.Errorf("login failed")

	user, err := l.dao.WithEmail(email)

	if err != nil {
		l.logger.Warn("DAO failed to query for user by email.", zap.Error(err),
			zap.String("user_email", email))

		return nil, loginFailure
	}

	if user == nil {
		l.logger.Debug("User not found.", zap.String("user_email", email))

		return nil, loginFailure
	}

	if user.KDF.Validate(password) {
		jwtResult, err := l.jwt.Generate(user)

		if err != nil {
			l.logger.Warn("Failed to generate JWT token for user.", zap.Error(err),
				zap.Uint("user_id", user.ID), zap.String("user_email", user.Email))

			return nil, loginFailure
		}

		result := newLoginResult(user, (*jwtResult).SignedToken(), (*jwtResult).Token(), (*jwtResult).Claims())

		return &result, nil
	} else {
		l.logger.Debug("Login KDF verification failed.", zap.Uint("user_id", user.ID),
			zap.String("user_email", user.Email))

		return nil, loginFailure
	}
}

func newLoginResult(user *models.User, signedToken *string, token *jwt.Token, claims *auth.JWTClaims) interfaces.LoginResult {
	return loginResultImpl{
		user:        user,
		signedToken: signedToken,
		token:       token,
		claims:      claims,
	}
}

type loginResultImpl struct {
	user        *models.User
	signedToken *string
	token       *jwt.Token
	claims      *auth.JWTClaims
}

func (l loginResultImpl) User() *models.User {
	return l.user
}

func (l loginResultImpl) SignedToken() *string {
	return l.signedToken
}

func (l loginResultImpl) Token() *jwt.Token {
	return l.token
}

func (l loginResultImpl) Claims() *auth.JWTClaims {
	return l.claims
}
