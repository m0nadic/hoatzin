package cmd

import (
	"fmt"
	"hoatzin/internal/app/server"
	"os"

	"github.com/spf13/cobra"
)

var bidiCmd = &cobra.Command{
	Use:   "bidi",
	Short: "Bi-directional communication with an interactive script.",
	Long: `Write received messages from the websocket to standard input
of an interactive program and stream the standard output to the websocket.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "ERROR:", "missing command")
			os.Exit(1)
		}
		bidiServer := server.NewBIDIServer(host, port, args)
		err := bidiServer.Start()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(bidiCmd)
}
