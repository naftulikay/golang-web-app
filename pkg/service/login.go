package service

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
)

type LoginServiceImpl struct {
	dao interfaces.UserDao
}

func NewLoginService(dao interfaces.UserDao) LoginServiceImpl {
	return LoginServiceImpl{dao: dao}
}

func (l LoginServiceImpl) Login(username, password string) (*models.User, error) {
	user, err := l.dao.WithUsername(username)

	if err != nil {
		return nil, fmt.Errorf("hidden login service error")
	}

	if user == nil {
		return nil, fmt.Errorf("login failed")
	}

	if user.KDF.Validate(password) {
		return user, nil
	} else {
		return nil, fmt.Errorf("login failed")
	}
}
