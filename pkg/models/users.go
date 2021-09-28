package models

import "gorm.io/gorm"

type User struct {
	Username string
	gorm.Model
}
