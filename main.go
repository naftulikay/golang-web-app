package main

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd"
)

// @title golang-webapp API
// @version v1
// @description Such a webapp, the most Go-est webapp.
// @BasePath /api/v1

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	if err := cmd.Execute(); err != nil {
		_ = fmt.Errorf("ERROR: failed to execute command: %s\n", err)
	}
}
