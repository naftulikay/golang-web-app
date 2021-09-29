package service

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/results"
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

		result := results.NewLoginResult(*user, (*jwtResult).SignedToken(), (*jwtResult).Token(), (*jwtResult).Claims())

		return &result, nil
	} else {
		l.logger.Debug("Login KDF verification failed.", zap.Uint("user_id", user.ID),
			zap.String("user_email", user.Email))

		return nil, loginFailure
	}
}
