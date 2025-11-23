package todo

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestAddTodo(t *testing.T) {
	tempFile := setupTestFile(t)

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

func TestDelete(t *testing.T) {
	tempFile := setupTestFile(t)

	_ = AddTodo("Buy milk")
	_ = AddTodo("Practice Vim")

	message, err := Delete("1")
	if err != nil {
		t.Fatalf("Delete returned an error: %v", err)
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
		t.Fatalf("expected 2 rows got %d", len(rows))
	}

	if rows[1][1] != "Practice Vim" {
		t.Fatalf("task different than expected: %s", rows[1][1])
	}

	if message != "Task deleted!" {
		t.Fatalf("message different than expected: %s", message)
	}
}

func TestComplete(t *testing.T) {
	tempFile := setupTestFile(t)

	_ = AddTodo("Practice Vim")

	message, err := Complete("1")
	if err != nil {
		t.Fatalf("Complete returned an error: %v", err)
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
		t.Fatalf("expected 2 rows got %d", len(rows))
	}

	if rows[1][2] != "Done" {
		t.Fatalf("status different than expected: %s", rows[1][1])
	}

	if message != "Task completed!" {
		t.Fatalf("message different than expected: %s", message)
	}

}

func setupTestFile(t *testing.T) string {
	t.Helper()

	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "todos.csv")

	getFilePath = func() (string, error) {
		return tempFile, nil
	}
	return tempFile
}
