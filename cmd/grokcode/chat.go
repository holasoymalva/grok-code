package grokcode

import (
	"fmt"
	"os"

	"github.com/holasoymalva/grok-code/internal/tui"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start the Grok Code interactive chat TUI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Grok Code...")
		if err := tui.RunTUI(); err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
