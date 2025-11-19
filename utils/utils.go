package utils

import (
	"os"
	"path/filepath"
)

func GetFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".todo-cli-app")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, "todos.csv"), nil
}
