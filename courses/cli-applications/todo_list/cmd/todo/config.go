package main

const (
	todosEnvFileName             = "TODOS_FILENAME"
	todosErrCannotReadFile       = 1
	todosErrCannotSaveFile       = 2
	todosErrCannotExecuteCommand = 3
)

var todosFilename = "todos.json"
