package views

import "time"

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserID    uint
	Email     string
	Username  string
	FirstName string
	LastName  string
	Role      string
	Token     string
	IssuedAt  time.Time
	NotBefore time.Time
	ExpiresAt time.Time
}
