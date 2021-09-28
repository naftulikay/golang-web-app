package interfaces

import "github.com/naftulikay/golang-webapp/pkg/models"

type UserDao interface {
	Get(id uint) (*models.User, error)
	WithUsername(username string) (*models.User, error)
}
