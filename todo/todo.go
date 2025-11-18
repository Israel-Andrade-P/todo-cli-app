package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type (
	TodoList struct {
		IdGenerator int
		Todos       []Todo
	}
	Todo struct {
		ID        string
		Task      string
		Status    string
		CreatedAt string
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

	writer.Write([]string{strconv.Itoa(id), body, "Not done", time.Now().Format("2006-01-02 15:04:05")})
	writer.Flush()
	return nil
}

func ListAll() error {
	path := "/home/zel/Go/todo-cli-app/todos.csv"
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

	todos, err := getTodos(path)
	if err != nil {
		return err
	}
	layout := "2006-01-02 15:04:05"
	for _, todo := range todos {
		loc := time.Now().Location()
		t, err := time.ParseInLocation(layout, todo.CreatedAt, loc)
		if err != nil {
			return err
		}
		writer.Write([]byte(fmt.Sprintf("%s\t%s\t%s\t%s\n", todo.ID, todo.Task, todo.Status, timediff.TimeDiff(t))))
	}
	defer writer.Flush()
	return nil
}

func Delete(id string) (string, error) {
	path := "/home/zel/Go/todo-cli-app/todos.csv"
	message := ""
	todos, err := getTodos(path)
	if err != nil {
		return message, err
	}
	deleted := false
	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			deleted = true
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return message, err
	}

	writer := csv.NewWriter(f)

	defer writer.Flush()

	writer.Write([]string{"ID", "Task", "Status", "Created"})
	for _, todo := range todos {
		writer.Write([]string{todo.ID, todo.Task, todo.Status, todo.CreatedAt})
	}
	if !deleted {
		message = fmt.Sprintf("Task with ID %s doesn't exist.", id)
		return message, nil
	}
	message = "Task deleted!"

	return message, nil
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

func getTodos(path string) ([]Todo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var todos []Todo
	for _, row := range rows[1:] {
		todos = append(todos, Todo{ID: row[0], Task: row[1], Status: row[2], CreatedAt: row[3]})
	}
	return todos, nil
}
