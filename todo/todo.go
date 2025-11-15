package todo

import "time"

type (
	TodoList struct {
		Todos []Todo
	}
	Todo struct {
		Task      string
		Status    bool
		CreatedAt time.Time
	}
)

func CreateTodoList() TodoList {
	return TodoList{Todos: make([]Todo, 0)}
}

func (todoList *TodoList) AddTodo(body string) {
	todoList.Todos = append(todoList.Todos, Todo{Task: body, Status: false, CreatedAt: time.Now()})
}
