package tools

import (
	"github.com/go-git/go-git/v5"
)

// GetStatus returns the current git status
func GetStatus(path string) (string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	w, err := r.Worktree()
	if err != nil {
		return "", err
	}

	status, err := w.Status()
	if err != nil {
		return "", err
	}

	return status.String(), nil
}
