package interfaces

import "github.com/naftulikay/golang-webapp/pkg/models"

type JWTService interface {
	Validate(token string) bool
}

type LoginService interface {
	Login(username, password string) (*models.User, error)
}
