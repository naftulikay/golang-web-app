package database

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	MaxDatabaseConnectionRequests = uint64(15)
)

type DatabaseLogger *zap.Logger

// Connect Attempt to connect to the database using an exponential backoff algorithm with a maximum retry count.
func Connect(config interfaces.MySQLConfig, dbLogger DatabaseLogger) (*gorm.DB, error) {
	logger := zap.Logger(*dbLogger)

	attempt := 0
	maxAttempts := MaxDatabaseConnectionRequests

	var db *gorm.DB

	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", config.User(),
		config.Password(), config.Host(), config.Port(), config.Database())
	sanitizedUri := strings.ReplaceAll(uri, config.Password(), "xxxxxxxxx")

	logger.Debug("Connecting to database.", zap.String("uri", sanitizedUri))

	connect := func() error {
		conn, err := gorm.Open(mysql.Open(uri))

		if err != nil {
			return err
		}

		// save the database connection
		db = conn

		return nil
	}

	notify := func(err error, next time.Duration) {
		attempt += 1
		logger.Debug(fmt.Sprintf("Connection attempt %d/%d failed, next retry in %s", attempt, maxAttempts,
			next.String()), zap.Error(err), zap.Duration("next_attempt", next))
	}

	err := backoff.RetryNotify(
		connect,
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxAttempts),
		notify,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("Unable to connect to database after %d attempts.", maxAttempts),
			zap.Error(err))

		return nil, fmt.Errorf("unable to connect to database after %d retries: %s",
			maxAttempts, err)
	}

	logger.Info("Connected to database.", zap.String("uri", sanitizedUri))

	return db, nil
}

// AutoMigrate Attempt to automatically migrate all database models.
func AutoMigrate(db *gorm.DB, logger *zap.Logger) error {
	logger.Debug("Running automatic database migrations...")
	err := db.AutoMigrate(models.Models()...)

	if err != nil {
		logger.Error("Unable to run automatic database migrations.", zap.Error(err))
	} else {
		logger.Info("Database migrated successfully.")
	}

	return err
}
