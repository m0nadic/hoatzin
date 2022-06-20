package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hoatzin/internal/app/server"
	"os"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Stream messages received as REST endpoint request",
	Long: `Starts a REST endpoint and streams the incoming request
body as websocket messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewRESTStreamingServer(host, port)
		err := s.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
