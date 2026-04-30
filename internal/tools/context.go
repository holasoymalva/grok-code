package tools

import (
	"os"
	"path/filepath"
	"strings"
)

// IndexRepo walks the directory and returns a simple tree representation
func IndexRepo(root string) (string, error) {
	var builder strings.Builder

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip hidden dirs like .git
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			return filepath.SkipDir
		}

		relPath, _ := filepath.Rel(root, path)
		if relPath == "." {
			return nil
		}

		builder.WriteString(relPath + "\n")
		return nil
	})

	return builder.String(), err
}
