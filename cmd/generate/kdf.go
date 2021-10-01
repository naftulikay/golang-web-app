package generate

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/naftulikay/golang-webapp/pkg/models"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	generateKDFCommand = &cobra.Command{
		Use:   "kdf",
		Short: "Generate a KDF object.",
		Run: func(cmd *cobra.Command, args []string) {
			password, err := cmd.Flags().GetString("password")

			if err != nil || len(password) == 0 {
				log.Fatalf("Unable to get password from CLI, pass --password|-p to set a password.")
			}

			kdf := models.GenKDF(password)

			output, _ := json.MarshalIndent(map[string]interface{}{
				"kdf_algorithm":     kdf.Algorithm,
				"kdf_password_hash": strings.ToUpper(hex.EncodeToString(kdf.PasswordHash[:])),
				"kdf_salt":          strings.ToUpper(hex.EncodeToString(kdf.Salt[:])),
				"kdf_time_factor":   kdf.TimeFactor,
				"kdf_memory_factor": kdf.MemoryFactor,
				"kdf_thread_factor": kdf.ThreadFactor,
				"kdf_key_len":       kdf.KeyLen,
			}, "", "  ")

			fmt.Println(string(output))
		},
	}
)

func init() {
	flags := generateKDFCommand.Flags()
	flags.StringP("password", "p", "", "The password to generate a KDF for.")
}
