package cmd

import (
	"fmt"
	"hoatzin/internal/app/server"
	"os"

	"github.com/spf13/cobra"
)

var (
	message  string
	interval int
	count    int
)
var tickCmd = &cobra.Command{
	Use:   "tick",
	Short: "Send message in every interval",
	Long:  `Executes a message template and sends it over websocket in every interval.`,
	Run: func(cmd *cobra.Command, args []string) {
		tts := server.NewTickingTemplateServer(host, port, message, interval, count)
		err := tts.Start()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tickCmd)

	tickCmd.Flags().IntVarP(&interval, "interval", "i", 1000, "Ticker interval in ms")
	tickCmd.Flags().StringVarP(&message, "message", "m", "", "Message template to be used. If this is provided, then rest of the args are ignored")
	tickCmd.Flags().IntVarP(&count, "count", "n", 1000, "Number of messages to send.")
}
