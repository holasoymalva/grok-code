package grokcode

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grokcode",
	Short: "Grok Code is an agentic coding assistant",
	Long:  `A multi-model CLI + TUI coding assistant agent written in Go, inspired by Claude Code but with "Grok vibes".`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no args, maybe start the TUI or show help
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("model", "", "Override the default model")
}
