package main

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_ = fmt.Errorf("ERROR: failed to execute command: %s\n", err)
	}
}
