package service

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"gorm.io/gorm"
	"testing"
)

type MockUserDao struct {
	returnError bool
	returnNil   bool
	interfaces.UserDao
}

func (m MockUserDao) WithEmail(email string) (*models.User, error) {
	if m.returnNil {
		return nil, nil
	}

	if m.returnError {
		return nil, fmt.Errorf("mock error")
	}

	return &models.User{
		Email:     email,
		FirstName: "Donaldus",
		LastName:  "Danguson",
		Role:      models.UserTypeNormal,
		KDF:       models.GenKDF("password"),
		Model: gorm.Model{
			ID: 123,
		},
	}, nil
}

type MockJWTService struct {
	returnError bool
	actual      JWTServiceImpl
	interfaces.JWTService
}

func (m MockJWTService) Generate(user *models.User) (interfaces.JWTGenerateResult, error) {
	if m.returnError {
		return nil, fmt.Errorf("mock error")
	}

	return m.actual.Generate(user)
}

func TestLoginServiceImpl_Interfaces(t *testing.T) {
	var _ interfaces.LoginService = LoginServiceImpl{}
	var _ interfaces.UserDao = MockUserDao{}
	var _ interfaces.JWTService = MockJWTService{}
}

func TestLoginServiceImpl_Login_Success(t *testing.T) {
	t.Skip("unimplemented")
}

func TestLoginServiceImpl_Login_DaoError(t *testing.T) {
	t.Skip("unimplemented")
}

func TestLoginServiceImpl_Login_NotFound(t *testing.T) {
	t.Skip("unimplemented")
}

func TestLoginServiceImpl_Login_WrongPassword(t *testing.T) {
	t.Skip("unimplemented")
}

func TestLoginServiceImpl_Login_JWTError(t *testing.T) {
	t.Skip("unimplemented")
}
