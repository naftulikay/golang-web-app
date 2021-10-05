package dao

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"github.com/naftulikay/golang-webapp/pkg/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserDao(db *gorm.DB, logger types.UserDaoLogger) *UserDaoImpl {
	return &UserDaoImpl{
		db:     db,
		logger: logger,
	}
}

func (u UserDaoImpl) Get(id uint) (*models.User, error) {
	var result models.User

	if u.db.First(&result, id).RowsAffected == 0 {
		return nil, fmt.Errorf("unable to find user with ID: %d", id)
	}

	return &result, nil
}

func (u UserDaoImpl) WithEmail(email string) (*models.User, error) {
	var result models.User

	if u.db.Where("email = ?", email).First(&result).RowsAffected == 0 {
		return nil, fmt.Errorf("unable to find user with email: %s", email)
	}

	return &result, nil
}
