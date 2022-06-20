package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to a topic and stream messages",
	Long: `Streams messages from a redis topic after
subscribing to the topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subscribe called")
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)
}
