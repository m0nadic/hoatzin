package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var interactCmd = &cobra.Command{
	Use:   "interact",
	Short: "Interact with websocket connection",
	Long: `Start an interactive shell to send and receive 
messages to and from the websocket connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("interact called")
	},
}

func init() {
	rootCmd.AddCommand(interactCmd)
}
