package service

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"testing"
)

func TestLoginServiceImpl_Interfaces(t *testing.T) {
	var _ interfaces.LoginService = LoginServiceImpl{}
}
