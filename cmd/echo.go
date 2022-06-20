package cmd

import (
	"fmt"
	"hoatzin/internal/app/server"
	"os"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echoes back the request",
	Long:  `Echoes back the websocket request back in the same connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		echoServer := server.NewEchoServer(host, port)
		err := echoServer.Start()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
