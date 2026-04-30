package agent

import (
	"context"

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

// RunLoop executes the agentic loop and returns the reply
func (a *Agent) RunLoop(ctx context.Context, task string) (string, error) {
	a.Memory = append(a.Memory, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: task,
	})

	req := openai.ChatCompletionRequest{
		Model:    a.Model,
		Messages: a.Memory,
	}

	resp, err := a.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	reply := resp.Choices[0].Message
	a.Memory = append(a.Memory, reply)

	return reply.Content, nil
}
