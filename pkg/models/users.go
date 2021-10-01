package models

import (
	"gorm.io/gorm"
	"strings"
)

const (
	UserTypeNormal = "user"
	UserTypeAdmin  = "admin"
)

type User struct {
	Email     string
	FirstName string
	LastName  string
	Role      string `gorm:"type:enum('user','admin')"`
	KDF       KDF    `gorm:"embedded;embeddedPrefix:kdf_"`
	gorm.Model
}

func (u User) Name() string {
	return strings.Join([]string{strings.TrimSpace(u.FirstName), strings.TrimSpace(u.LastName)}, " ")
}
