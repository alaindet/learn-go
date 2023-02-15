package main

import (
	"fmt"
	"os"

	"todo_list/internal/todo"
)

func addTodoCommand(todos *todo.Todos, newTodo string) {
	todos.Add(newTodo)
	err := todos.SaveToStorage(todosFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotSaveFile)
	}
	fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
	showListCommand(todos)
}

func showListCommand(todos *todo.Todos) {
	for _, todo := range *todos {

		checkbox := "[ ]"
		if todo.Done {
			checkbox = "[x]"
		}

		fmt.Printf("%s %s\n", checkbox, todo.Name)
	}
}

func completeTodoCommand(todos *todo.Todos, todoIndex int) {
	todos.Complete(todoIndex)
	err := todos.SaveToStorage(todosFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(todosErrCannotSaveFile)
	}
	fmt.Printf("Todo #%d completed.\n\n", todoIndex)
	showListCommand(todos)
}
