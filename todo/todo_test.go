package todo

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestAddTodo(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "todos.csv")

	getFilePath = func() (string, error) {
		return tempFile, nil
	}

	err := AddTodo("Buy milk")
	if err != nil {
		t.Fatalf("AddTodo returned an error: %v", err)
	}

	f, err := os.Open(tempFile)
	if err != nil {
		t.Fatalf("failed to open created csv file: %v", err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("failed to read csv file: %v", err)
	}

	if len(rows) != 2 {
		t.Fatalf("expected two rows, got %d", len(rows))
	}

	header := rows[0]
	if header[0] != "ID" || header[1] != "Task" {
		t.Fatalf("header row incorrect, %v", header)
	}

	todoRow := rows[1]
	if todoRow[1] != "Buy milk" {
		t.Fatalf("expected task 'Buy milk' got %s", todoRow[1])
	}

	if todoRow[2] != "pending" {
		t.Fatalf("expected status 'pending' got %s", todoRow[2])
	}
}
