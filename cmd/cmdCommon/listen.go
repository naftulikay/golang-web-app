package cmdCommon

import (
	"fmt"
	"github.com/naftulikay/golang-webapp/cmd/cmdConstants"
	"github.com/spf13/pflag"
	"strings"
)

func ListenFlags(flags *pflag.FlagSet) {
	// --listen
	flags.StringP(cmdConstants.CliFlagListen, "H", cmdConstants.DefaultListenHost,
		fmt.Sprintf("The host to listen on for incoming connections. [env: %s]",
			strings.ToUpper(cmdConstants.EnvVarListenHost)))
	// --port
	flags.Uint16P(cmdConstants.CliFlagPort, "p", cmdConstants.DefaultListenPort,
		fmt.Sprintf("The port to listen on for incoming connections. [env: %s]",
			strings.ToUpper(cmdConstants.CliFlagPort)))
}
