package dao

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"gorm.io/gorm"
)

type UserDaoImpl struct {
	DB *gorm.DB
}

func (u UserDaoImpl) Get(id uint) (*models.User, error) {
	var result models.User

	if u.DB.First(&result, id).RowsAffected == 0 {
		return nil, fmt.Errorf("unable to find user with ID: %d", id)
	}

	return &result, nil
}

func (u UserDaoImpl) WithEmail(username string) (*models.User, error) {
	var result models.User

	if u.DB.Where("username = ?", username).First(&result).RowsAffected == 0 {
		return nil, fmt.Errorf("unable to find user with username: %s", username)
	}

	return &result, nil
}
