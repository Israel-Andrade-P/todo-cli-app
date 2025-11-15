package todo

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type (
	TodoList struct {
		IdGenerator int
		Todos       []Todo
	}
	Todo struct {
		ID        int
		Task      string
		Status    bool
		CreatedAt time.Time
	}
)

func CreateTodoList() TodoList {
	return TodoList{Todos: make([]Todo, 0)}
}

func AddTodo(body string) error {
	path := "/home/zel/Go/todo-cli-app/todos.csv"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(f)

	if info.Size() == 0 {
		writer.Write([]string{"ID", "Task", "Status", "Created"})
		writer.Flush()
	}

	id, err := getId(path)
	if err != nil {
		return err
	}

	writer.Write([]string{strconv.Itoa(id), body, "Not done", "Now"})
	writer.Flush()
	return nil
}

func getId(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	rows, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}
	return len(rows), nil
}
