package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ProviderConfig struct {
	BaseURL string `yaml:"base_url"`
	APIKey  string `yaml:"api_key"`
}

type CostEstimator struct {
	WarningThreshold float64 `yaml:"warning_threshold"`
}

type Config struct {
	DefaultModel  string                    `yaml:"default_model"`
	Providers     map[string]ProviderConfig `yaml:"providers"`
	CostEstimator CostEstimator             `yaml:"cost_estimator"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Router helps pick the best provider
func (c *Config) GetProvider(modelMode string) ProviderConfig {
	// Simple routing logic: cheap mode, reasoning mode, fast mode
	switch modelMode {
	case "cheap":
		return c.Providers["deepseek"]
	case "reasoning":
		return c.Providers["xai"]
	case "fast":
		return c.Providers["gemini"]
	default:
		// Try default
		if p, ok := c.Providers[c.DefaultModel]; ok {
			return p
		}
		// Fallback
		return c.Providers["xai"]
	}
}
