package cmd

import (
	"file-transfer/internal/server"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	port       int
	storageDir string
)

var rootCmd = &cobra.Command{
	Use:   "serveme",
	Short: "Start a storage server",
	Long: `CLI command to start a web server that accepts files via HTTP and stores them.
You can retrieve them too`,
	Run: run,
}

func run(_ *cobra.Command, _ []string) {
	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		slog.Error("internal server error: ", err)
	}

	app := app.New(port, storageDir)
	fmt.Printf("Server is starting on http://localhost:%d", port)
	err = app.Run()
	if err != nil {
		slog.Error("internal server error: ", err)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Specifies on what port the server runs")
	rootCmd.PersistentFlags().StringVarP(&storageDir, "storage", "s", "storage", "Specifies where files are stored")
}

func Execute() {
	rootCmd.Execute()
}
