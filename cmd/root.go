package cmd

import (
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/naftulikay/golang-webapp/cmd/serve"
	"github.com/naftulikay/golang-webapp/cmd/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	rootCommand = &cobra.Command{
		Use:   "golang-webapp",
		Short: "golang-webapp application.",
	}
)

func Execute() error {
	return rootCommand.Execute()
}

func init() {
	rootCommand.AddCommand(serve.Commands()...)
	rootCommand.AddCommand(service.Commands()...)

	// global viper initialization
	viper.AutomaticEnv()

	// register all known environment variables
	for _, v := range cmdConstants.EnvVariables() {
		if err := viper.BindEnv(v); err != nil {
			log.Fatalf("Unable to bind environment variable %s: %s", strings.ToUpper(v), err)
		}
	}
}
