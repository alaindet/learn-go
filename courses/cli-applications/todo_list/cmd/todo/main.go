package main

import (
	"fmt"
	"os"

	"todo_list/internal/todo"
)

func main() {
	input, todos := setup()
	execute(input, todos)
}

func setup() (*TodoCliInput, *todo.Todos) {
	input := NewInput()
	input.Parse()

	todos := todo.NewTodos()
	err := todos.FetchFromStorage(todosFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotReadFile)
	}

	return input, todos
}

func execute(input *TodoCliInput, todos *todo.Todos) {
	switch {
	case input.addTodo.Value() != "":
		addTodoCommand(todos, input.addTodo.Value())
	case input.showList.Value():
		showListCommand(todos)
	case input.completeTodo.Value() != -1:
		completeTodoCommand(todos, input.completeTodo.Value())
	default:
		showListCommand(todos)
	}
}
