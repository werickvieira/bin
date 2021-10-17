package cmd

import (
	"log"

	"bin/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt is used to decrypt the file",
	Long:  `Decrypt is used to call the function that should undo the cryptography in the file`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := viper.GetString("file.name")
		fileExtension := viper.GetString("file.extension")
		key := viper.GetString("key")

		if err := app.DecryptFile(fileName+fileExtension, key); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.PersistentFlags().StringP("decrypt", "d", "decrypt", "undos the cryptography")
}
