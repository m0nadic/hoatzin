package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hoatzin/internal/app/server"
	"os"
)

var (
	redisHost    string
	redisPort    int
	redisChannel string
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to a topic and stream messages",
	Long: `Streams messages from a redis topic after
subscribing to the topic.`,
	Run: func(cmd *cobra.Command, args []string) {
		ss := server.NewSubscribingServer(host, port, redisHost, redisPort, redisChannel)
		err := ss.Start()
		
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)
	subscribeCmd.Flags().StringVar(&redisHost, "redis-host", "localhost", "Redis host name or IP.")
	subscribeCmd.Flags().IntVar(&redisPort, "redis-port", 6379, "Redis port.")
	subscribeCmd.Flags().StringVar(&redisChannel, "channel", "hoatzin", "Redis channel to subscribe.")

}
