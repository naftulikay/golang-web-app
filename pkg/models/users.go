package models

import "gorm.io/gorm"

type User struct {
	Username string
	KDF      KDF `gorm:"embedded;embeddedPrefix:kdf_"`
	gorm.Model
}
