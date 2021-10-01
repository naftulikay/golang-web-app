package generate

import "github.com/spf13/cobra"

var (
	generateUserCommand = &cobra.Command{
		Use:   "user",
		Short: "Generate a user object for insertion into the database.",
		PreRun: func(cmd *cobra.Command, args []string) {
			panic("unimplemented!")
		},
		Run: func(cmd *cobra.Command, args []string) {
			panic("unimplemented!")
		},
	}
)

func init() {

}
