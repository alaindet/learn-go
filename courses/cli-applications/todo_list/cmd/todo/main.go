package main

import (
	"fmt"
)

const (
	todosFilename          = "todos.json"
	todosErrCannotReadFile = 1
	todosErrCannotSaveFile = 2
)

func main() {

	input := NewInput()
	input.Parse()

	fmt.Printf("addTodo: (%T) %+v\n", input.addTodo.value, input.addTodo.value)
	fmt.Printf("showList: (%T) %+v\n", input.showList.value, input.showList.value)
	fmt.Printf("completeTodo: (%T) %+v\n", input.completeTodo.value, input.completeTodo.value)

	// todos := todo.NewTodos()
	// err := todos.FetchFromStorage(todosFilename)

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(todosErrCannotReadFile)
	// }

	// // No args, list all todo items
	// if len(os.Args) == 1 {
	// 	for _, todo := range *todos {
	// 		fmt.Println(todo.Name)
	// 	}
	// 	return
	// }

	// // Merge all input as a new todo item
	// item := strings.Join(os.Args[1:], " ")
	// todos.Add(item)
	// err = todos.SaveToStorage(todosFilename)

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(todosErrCannotSaveFile)
	// }
}
