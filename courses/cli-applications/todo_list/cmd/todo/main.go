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

	if os.Getenv(todosEnvFileName) != "" {
		todosFilename = os.Getenv(todosEnvFileName)
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
	case input.addTodo:
		err = addTodoCommand(todos)
	case input.showList:
		err = showListCommand(todos)
	case input.completeTodo != -1:
		err = completeTodoCommand(todos, input.completeTodo)
	default:
		err = showListCommand(todos)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotExecuteCommand)
	}
}
