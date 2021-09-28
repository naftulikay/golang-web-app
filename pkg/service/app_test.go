package service

import (
	"github.com/naftulikay/golang-webapp/pkg/interfaces"
	"testing"
)

// TestAppImpl_Interfaces Tests that AppImpl implements the interfaces.App interface.
func TestAppImpl_Interfaces(t *testing.T) {
	var _ interfaces.App = AppImpl{}
}

// TestAppServicesImpl_Interfaces Tests that AppServicesImpl implements the interfaces.AppServices interface.
func TestAppServicesImpl_Interfaces(t *testing.T) {
	var _ interfaces.AppServices = AppServicesImpl{}
}

// TestAppDaosImpl_Interfaces Tests that AppDaos implements the interfaces.AppDaos interface
func TestAppDaosImpl_Interfaces(t *testing.T) {
	var _ interfaces.AppDaos = AppDaosImpl{}
}
