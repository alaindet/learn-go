package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"todo_list/internal/todo"
	"todo_list/internal/try"
)

func addTodoCommand(todos *todo.Todos) error {
	return try.Try(func(catch try.Catcher) {
		newTodo, err := getTask(os.Stdin, flag.Args()...)
		catch(err)
		catch(todos.Add(newTodo))
		catch(todos.SaveToStorage(todosFilename))
		fmt.Printf("Todo \"%s\" created.\n\n", newTodo)
		catch(showListCommand(todos))
	})
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	err := s.Err()
	if err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("new todo cannot be blank")
	}
	return s.Text(), nil
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
