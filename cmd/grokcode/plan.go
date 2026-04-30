package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan [description]",
	Short: "Ask the agent to generate a plan for a given task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		fmt.Printf("Grok Code: Drafting a master plan for: %s\n", description)
		// Initialize the agent and run only the planning phase
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
