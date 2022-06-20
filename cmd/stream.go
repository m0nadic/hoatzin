package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hoatzin/internal/app/server"
	"os"
)

var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Streams the standard output.",
	Long:  `Streams the standard output of an executable program through a websocket.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "ERROR:", "missing command")
			os.Exit(1)
		}
		streamingServer := server.NewStreamingServer(host, port, args)
		err := streamingServer.Start()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
}
