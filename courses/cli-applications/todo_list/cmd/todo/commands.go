package main

import (
	"fmt"

	"todo_list/internal/todo"
	"todo_list/internal/try"
)

func addTodoCommand(todos *todo.Todos) error {
	return try.Try(func(err try.Catcher) {
		newTodo := "" // TODO
		err(todos.Add(newTodo))
		err(todos.SaveToStorage(todosFilename))
		fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
		err(showListCommand(todos))
	})
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
	return try.Try(func(err try.Catcher) {
		err(todos.Complete(todoIndex))
		err(todos.SaveToStorage(todosFilename))
		fmt.Printf("Todo #%d completed.\n\n", todoIndex)
		err(showListCommand(todos))
	})
}
