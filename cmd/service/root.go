package service

import "github.com/spf13/cobra"

var (
	serviceCommand = &cobra.Command{
		Use:   "service",
		Short: "Interact with service objects.",
	}
)

func Commands() []*cobra.Command {
	return []*cobra.Command{serviceCommand}
}

func init() {
	serviceCommand.AddCommand(loginCommand)
}
