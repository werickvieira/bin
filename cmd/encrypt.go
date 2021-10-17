package cmd

import (
	"bin/app"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt is used to encrypt the file.",
	Long:  `Encrypt is used to call the function that should make the cryptography in the file`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := viper.GetString("file.name")
		fileExtension := viper.GetString("file.extension")
		key := viper.GetString("key")

		if err := app.EncryptFile(fileName+fileExtension, key); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.PersistentFlags().StringP("encrypt", "e", "encrypt", "makes the cryptography")
}
