package interfaces

import "github.com/naftulikay/golang-webapp/pkg/models"

type UserDao interface {
	Get(id uint) (*models.User, error)
	Exists(id uint) (bool, error)
	WithEmail(username string) (*models.User, error)
}
