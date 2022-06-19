package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "hoatzin",
	Short: "A convenient websocket server.",
	Long:  `A general purpose websocket server which can be used to mock websocket backends.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "b", "0.0.0.0", "Address to bind to")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 4321, "HTTP port to listen on")
}
