package todo

import (
	"encoding/csv"
	"os"
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

func (todoList *TodoList) AddTodo(body string) error {
	path := "/home/zel/Go/todo-cli-app/todos.csv"
	todoList.IdGenerator++
	todoList.Todos = append(todoList.Todos, Todo{ID: todoList.IdGenerator, Task: body, Status: false, CreatedAt: time.Now()})
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
	defer writer.Flush()

	if info.Size() == 0 {
		writer.Write([]string{"ID", "Task", "Status", "Created"})
	}

	writer.Write([]string{"2", body, "Not done", "Now"})
	return nil
}
