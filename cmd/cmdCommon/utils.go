package cmdCommon

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/howeyc/gopass"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

// Fatal Log a fatal message and exit, using the zap.Logger if provided or log if not.
func Fatal(msg string, logger *zap.Logger) {
	if logger != nil {
		logger.Fatal(msg)
	} else {
		log.Fatal(msg)
	}
}

// StdinPassword Read a password from standard input, fatally exiting if a problem is encountered.
func StdinPassword(prompt string, logger *zap.Logger) string {
	if strings.HasSuffix(prompt, " ") {
		print(prompt)
	} else {
		print(prompt + " ")
	}

	pwbytes, err := gopass.GetPasswd()

	if err != nil {
		Fatal(fmt.Sprintf("Unable to read password from standard input: %s", err), logger)
	}

	if len(pwbytes) == 0 {
		Fatal("Please enter a password", logger)
	}

	return strings.TrimSpace(string(pwbytes))
}

func ValidateConfig(config interface{}, logger *zap.Logger) {
	v := validator.New()
	err := v.Struct(config)

	if err != nil {
		if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
			Fatal(fmt.Sprintf("Unable to validate configuration: %s", invalidErr), logger)
		} else if errors, ok := err.(validator.ValidationErrors); ok {
			var resultErr error

			for _, e := range errors {
				resultErr = multierror.Append(resultErr, e)
			}

			Fatal(fmt.Sprintf("Configuration is invalid: %s", resultErr), logger)
		}
	}
}

func Logger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()

	if err != nil {
		Fatal(fmt.Sprintf("Unable to create logger: %s", err), logger)
	}

	return logger
}
