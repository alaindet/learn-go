package main

import (
	"fmt"

	"todo_list/internal/todo"
)

func addTodoCommand(todos *todo.Todos) error {

	/*
		errors.Try(
			func (try errors.ErrorCatcher) {
				try(todos.Add(newTodo))
				try(todos.SaveToStorage(todosFilename))
				fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
				try(showListCommand(todos))
			},
			func (err error) {
				fmt.Println("Oops", err)
			}
		)

		errors.Try(
			func(try errors.ErrorCatcher) {
				try(todos.Add(newTodo))
				try(todos.SaveToStorage(todosFilename))
				fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
				try(showListCommand(todos))
			},
			func (err error) {
				fmt.Println("Oops", err)
			}
		)
	*/

	newTodo := "" // TODO

	err := todos.Add(newTodo)

	if err != nil {
		return err
	}

	err = todos.SaveToStorage(todosFilename)

	if err != nil {
		return err
	}

	fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
	err = showListCommand(todos)

	if err != nil {
		return err
	}

	return nil
}

func showListCommand(todos *todo.Todos) error {
	for _, todo := range *todos {

		checkbox := "[ ]"
		if todo.Done {
			checkbox = "[x]"
		}

		fmt.Printf("%s %s\n", checkbox, todo.Name)
	}

	return nil
}

func completeTodoCommand(todos *todo.Todos, todoIndex int) error {
	err := todos.Complete(todoIndex)

	if err != nil {
		return err
	}

	storeTodos(todos)
	fmt.Printf("Todo #%d completed.\n\n", todoIndex)
	showListCommand(todos)

	return nil
}
