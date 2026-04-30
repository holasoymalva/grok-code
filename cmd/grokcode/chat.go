package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/holasoymalva/grok-code/internal/agent"
	"github.com/holasoymalva/grok-code/internal/config"
	"github.com/holasoymalva/grok-code/internal/tui"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start the Grok Code interactive chat TUI",
	Run: func(cmd *cobra.Command, args []string) {
		// Intro Animation
		banner := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00ADD8")).
			Render(`
   _____           _       _____          _      
  / ____|         | |     / ____|        | |     
 | |  __ _ __ ___ | | __ | |     ___   __| | ___ 
 | | |_ | '__/ _ \| |/ / | |    / _ \ / _' |/ _ \
 | |__| | | | (_) |   <  | |___| (_) | (_| |  __/
  \_____|_|  \___/|_|\_\  \_____\___/ \__,_|\___|
`)
		fmt.Println(banner)
		fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Initializing Agentic Environment..."))
		time.Sleep(1200 * time.Millisecond)

		cfg, err := config.LoadConfig("config.yaml")
		
		var provider config.ProviderConfig
		var modelName string

		if err != nil {
			fmt.Println("Info: No config.yaml found. Using free public endpoint (Pollinations.ai).")
			provider = config.ProviderConfig{
				BaseURL: "https://text.pollinations.ai/openai/v1",
				APIKey:  "anonymous",
			}
			modelName = "openai"
		} else {
			provider = cfg.GetProvider("fast")
			modelName = "gemini-2.5-flash"
			
			if provider.APIKey == "" {
				fmt.Println("Info: API Key is empty. Falling back to free public endpoint (Pollinations.ai).")
				provider = config.ProviderConfig{
					BaseURL: "https://text.pollinations.ai/openai/v1",
					APIKey:  "anonymous",
				}
				modelName = "openai"
			}
		}

		ag := agent.NewAgent(provider.APIKey, provider.BaseURL, modelName)

		if err := tui.RunTUI(ag); err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
