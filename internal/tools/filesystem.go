package tools

import (
	"os"
)

// ReadFile reads the content of a file
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteFile writes content to a file, potentially showing a diff first
func WriteFile(path string, content string) error {
	// In a real implementation, we would show a diff here before writing
	return os.WriteFile(path, []byte(content), 0644)
}
