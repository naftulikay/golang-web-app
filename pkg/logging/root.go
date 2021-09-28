package logging

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/constants"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Configure Setup zap logging and return the root logger.
func Configure(config interfaces.AppConfig) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	switch env := config.Env(); env {
	case constants.DevEnvironment:
		// for development, use the human-readable logger
		logger, err = configureDev(config)
	case constants.StagingEnvironment:
		fallthrough
	case constants.ProductionEnvironment:
		// for staging and production, use a production logger
		logger, err = configureProd(config)
	default:
		return nil, fmt.Errorf("unknown execution environment '%s', acceptable environments are: %v",
			config.Env(), constants.EnvironmentList())
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func configureDev(appConfig interfaces.AppConfig) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return config.Build()
}

func configureProd(config interfaces.AppConfig) (*zap.Logger, error) {
	return zap.NewProduction()
}
