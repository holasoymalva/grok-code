package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/holasoymalva/grok-code/internal/tools"
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

	toolsDef := []openai.Tool{
		{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "write_file",
				Description: "Write content to a file at the specified path. Creates folders if they don't exist.",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "The absolute or relative path to the file",
						},
						"content": map[string]interface{}{
							"type":        "string",
							"description": "The full content to write to the file",
						},
					},
					"required": []string{"path", "content"},
				},
			},
		},
		{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        "read_file",
				Description: "Read the content of a file",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"path": map[string]interface{}{
							"type":        "string",
							"description": "The absolute or relative path to the file",
						},
					},
					"required": []string{"path"},
				},
			},
		},
	}

	for {
		req := openai.ChatCompletionRequest{
			Model:    a.Model,
			Messages: a.Memory,
			Tools:    toolsDef,
		}

		resp, err := a.Client.CreateChatCompletion(ctx, req)
		if err != nil {
			return "", err
		}

		reply := resp.Choices[0].Message
		a.Memory = append(a.Memory, reply)

		if len(reply.ToolCalls) == 0 {
			// No more tool calls, return final response
			return reply.Content, nil
		}

		// Handle tool calls
		for _, tc := range reply.ToolCalls {
			if tc.Type == openai.ToolTypeFunction {
				var resultStr string

				if tc.Function.Name == "write_file" {
					var args struct {
						Path    string `json:"path"`
						Content string `json:"content"`
					}
					if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err == nil {
						err := tools.WriteFile(args.Path, args.Content)
						if err != nil {
							resultStr = fmt.Sprintf("Error writing file: %v", err)
						} else {
							resultStr = fmt.Sprintf("Success: File %s was written.", args.Path)
						}
					} else {
						resultStr = "Error parsing arguments: " + err.Error()
					}
				} else if tc.Function.Name == "read_file" {
					var args struct {
						Path string `json:"path"`
					}
					if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err == nil {
						content, err := tools.ReadFile(args.Path)
						if err != nil {
							resultStr = fmt.Sprintf("Error reading file: %v", err)
						} else {
							resultStr = content
						}
					} else {
						resultStr = "Error parsing arguments: " + err.Error()
					}
				} else {
					resultStr = "Unknown function"
				}

				a.Memory = append(a.Memory, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					Content:    resultStr,
					Name:       tc.Function.Name,
					ToolCallID: tc.ID,
				})
			}
		}
	}
}
