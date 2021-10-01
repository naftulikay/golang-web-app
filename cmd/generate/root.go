package generate

import "github.com/spf13/cobra"

var (
	generateCommand = &cobra.Command{
		Use:   "generate",
		Short: "Generate values.",
	}
)

func Commands() []*cobra.Command {
	return []*cobra.Command{generateCommand}
}

func init() {
	generateCommand.AddCommand(generateUserCommand, generateKDFCommand)
}
