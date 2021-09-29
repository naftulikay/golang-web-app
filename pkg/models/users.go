package models

import "gorm.io/gorm"

const (
	UserTypeNormal = "user"
	UserTypeAdmin  = "admin"
)

type User struct {
	Email     string
	FirstName string
	LastName  string
	Role      string `gorm:"type('user','admin')"`
	KDF       KDF    `gorm:"embedded;embeddedPrefix:kdf_"`
	gorm.Model
}
