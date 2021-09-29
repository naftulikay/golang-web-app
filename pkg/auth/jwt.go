package auth

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	UserID    uint64 `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}
