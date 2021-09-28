package views

import "time"

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	UserID    uint
	Email     string
	Username  string
	FirstName string
	LastName  string
	Roles     []string
	Token     string
	IssuedAt  time.Time
	NotBefore time.Time
	ExpiresAt time.Time
}
