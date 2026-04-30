package agent

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	Client *openai.Client
	Model  string
	Memory []openai.ChatCompletionMessage
}

func NewAgent(apiKey string, baseURL string, model string) *Agent {
	cfg := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		cfg.BaseURL = baseURL
	}
	client := openai.NewClientWithConfig(cfg)

	return &Agent{
		Client: client,
		Model:  model,
		Memory: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: SystemPrompt,
			},
		},
	}
}

// RunLoop executes the agentic loop
func (a *Agent) RunLoop(ctx context.Context, task string) error {
	a.Memory = append(a.Memory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: task,
	})

	// Simplistic loop: Plan -> Execute -> Verify
	fmt.Println("Grok Code: Thinking...")

	// Normally this would use streaming and tool calls
	req := openai.ChatCompletionRequest{
		Model:    a.Model,
		Messages: a.Memory,
		// Tools: (Add tool schemas here)
	}

	resp, err := a.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return err
	}

	reply := resp.Choices[0].Message
	a.Memory = append(a.Memory, reply)

	fmt.Printf("Grok Code: %s\n", reply.Content)

	// Here we would implement the verification and self-correction loop
	// If reply has tool calls, execute them and feedback results.

	return nil
}
