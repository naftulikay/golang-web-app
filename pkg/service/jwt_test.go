package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"github.com/naftulikay/golang-webapp/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestJWTServiceImpl_Interfaces(t *testing.T) {
	var _ interfaces.JWTService = JWTServiceImpl{}
}

func TestJWTServiceImpl_Generate_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic with a zero key")
		}
	}()

	service := JWTServiceImpl{
		key: utils.NullBytes32(),
	}

	user := models.User{}

	_, _ = service.Generate(&user)
}

func TestJWTServiceImpl_Validate_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic with a zero key")
		}
	}()

	service := JWTServiceImpl{
		key: utils.NullBytes32(),
	}

	tokenString := utils.InfallibleSecureRandBase64(64)

	_, _ = service.Validate(tokenString)
}

func TestJWTServiceImpl_E2E(t *testing.T) {
	service := JWTServiceImpl{
		key: utils.InfallibleSecureRandBytes32(),
	}

	user := models.User{
		Email:     "donny.dangus@gmail.com",
		FirstName: "Donaldus",
		LastName:  "Dangus",
		Role:      models.UserTypeNormal,
		Model: gorm.Model{
			ID: 123,
		},
	}

	generated, err := service.Generate(&user)

	if err != nil {
		t.Errorf("unable to generate JWT token for user: %s", err)
	}

	assert.NotNil(t, generated)

	signed := (*generated).SignedToken()
	token := (*generated).Token()
	claims := (*generated).Claims()

	now := time.Now().UTC()

	assert.NotNil(t, signed)
	assert.NotNil(t, token)
	assert.NotNil(t, claims)

	assert.Equal(t, "HS256", token.Header["alg"])
	assert.Equal(t, jwt.SigningMethodHS256, token.Method)

	assert.Equal(t, uint64(user.ID), claims.UserID)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.FirstName, claims.FirstName)
	assert.Equal(t, user.LastName, claims.LastName)
	assert.Equal(t, user.Role, claims.Role)
	assert.Equal(t, JWTIssuer, claims.Issuer)
	assert.GreaterOrEqual(t, now.Unix(), claims.NotBefore)
	assert.GreaterOrEqual(t, now.Unix(), claims.IssuedAt)
	assert.Equal(t, claims.NotBefore+int64(JWTExpiry.Seconds()), claims.ExpiresAt)

	assert.Greater(t, len(*signed), 0)

	// generation testing is complete, now onto verification
	validated, err := service.Validate(*signed)

	assert.NotNil(t, validated)

	token = (*validated).Token()
	claims = (*validated).Claims()

	assert.NotNil(t, token)
	assert.NotNil(t, claims)
	assert.Nil(t, err)

	assert.True(t, token.Valid)
	assert.Equal(t, jwt.SigningMethodHS256, token.Method)

	assert.Equal(t, uint64(user.ID), claims.UserID)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.FirstName, claims.FirstName)
	assert.Equal(t, user.LastName, claims.LastName)
	assert.Equal(t, user.Role, claims.Role)

	assert.Equal(t, JWTIssuer, claims.Issuer)
	assert.GreaterOrEqual(t, now.Unix(), claims.NotBefore)
	assert.GreaterOrEqual(t, now.Unix(), claims.IssuedAt)
	assert.Equal(t, claims.NotBefore+int64(JWTExpiry.Seconds()), claims.ExpiresAt)

	ss, err := token.SignedString(service.key[:])

	assert.Nil(t, err)
	assert.Equal(t, *signed, ss)
}
