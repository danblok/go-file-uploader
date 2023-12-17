package cmd

import (
	upload "file-transfer/internal/client"
	"fmt"

	"github.com/spf13/cobra"
)

var to string

var rootCmd = &cobra.Command{
	Use:   "sendme",
	Short: "Uploads files via HTTP",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	uploader := upload.New(to)
	err := uploader.UploadFiles(args)
	if err != nil {
		fmt.Println("An error occured uploading the file(s)")
		return
	}
	fmt.Println("Successfully sent the file(s)!")
}

func init() {
	rootCmd.PersistentFlags().StringVar(&to, "to", "http://localhost:8080", "Sets host to send attached files to")
}

func Execute() {
	rootCmd.Execute()
}
