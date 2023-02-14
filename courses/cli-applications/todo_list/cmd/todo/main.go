package main

import (
	"fmt"
	"os"
	"strings"

	"todo_list/internal/todo"
)

const (
	todosFilename          = "todos.json"
	todosErrCannotReadFile = 1
	todosErrCannotSaveFile = 2
)

func main() {
	todos := todo.NewTodos()
	err := todos.FetchFromStorage(todosFilename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotReadFile)
	}

	// No args, list all todo items
	if len(os.Args) == 1 {
		for _, todo := range *todos {
			fmt.Println(todo.Name)
		}
		return
	}

	// Merge all input as a new todo item
	item := strings.Join(os.Args[1:], " ")
	todos.Add(item)
	err = todos.SaveToStorage(todosFilename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotSaveFile)
	}
}
