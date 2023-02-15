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

	if os.Getenv("TODOS_FILENAME") != "" {
		todosFilename = os.Getenv("TODOS_FILENAME")
	}

	todos := todo.NewTodos()
	err := todos.FetchFromStorage(todosFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotReadFile)
	}

	return input, todos
}

func execute(input *TodoCliInput, todos *todo.Todos) {
	var err error

	switch {
	case input.addTodo.Value():
		err = addTodoCommand(todos)
	case input.showList.Value():
		err = showListCommand(todos)
	case input.completeTodo.Value() != -1:
		err = completeTodoCommand(todos, input.completeTodo.Value())
	default:
		err = showListCommand(todos)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotExecuteCommand)
	}
}
