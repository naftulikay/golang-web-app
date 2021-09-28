package models

import "gorm.io/gorm"

type User struct {
	Username string
	KDF
	gorm.Model
}
