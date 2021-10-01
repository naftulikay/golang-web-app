package types

import (
	"go.uber.org/zap"
)

type RootLogger *zap.Logger
type AppLogger *zap.Logger
type JWTServiceLogger *zap.Logger
type LoginServiceLogger *zap.Logger
type UserDaoLogger *zap.Logger
type DatabaseLogger *zap.Logger

func NewLoginServiceLogger(logger AppLogger) LoginServiceLogger {
	return (*logger).Named("services.login")
}

func NewJWTServiceLogger(logger AppLogger) JWTServiceLogger {
	return (*logger).Named("services.jwt")
}

func NewUserDaoLogger(logger AppLogger) UserDaoLogger {
	return (*logger).Named("dao.users")
}

func NewDatabaseLogger(logger AppLogger) DatabaseLogger {
	return (*logger).Named("db")
}

func NewAppLogger(logger RootLogger) AppLogger {
	return (*logger).Named("app")
}
