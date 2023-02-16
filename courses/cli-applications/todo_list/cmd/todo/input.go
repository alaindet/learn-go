package main

import (
	"flag"
	"todo_list/internal/input"
)

type TodoCliInput struct {
	addTodo      bool
	showList     bool
	completeTodo int
}

func NewInput() *TodoCliInput {

	f := input.NewCLIFlags()

	f.Bool("add", false, "Add new todo from stdin (if provided) or args")
	f.Bool("list", false, "Show list of all todos")
	f.Int("complete", -1, "Complete given todo by index")

	f.Parse()

	return &TodoCliInput{
		addTodo:      f.GetBool("add"),
		showList:     f.GetBool("list"),
		completeTodo: f.GetInt("complete"),
	}
}

func (t *TodoCliInput) Parse() {
	flag.Parse()
}
