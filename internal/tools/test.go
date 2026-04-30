package tools

import (
	"strings"
)

// RunTests executes tests depending on the project type
func RunTests(projectType string) (string, error) {
	switch strings.ToLower(projectType) {
	case "go":
		return RunCommand("go", "test", "./...")
	case "npm":
		return RunCommand("npm", "test")
	default:
		return "", nil
	}
}
