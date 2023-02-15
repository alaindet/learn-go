package main

import "flag"

type InputFlag[T any] struct {
	name         string
	defaultValue T
	description  string
	value        *T
}

func NewInputFlag[T any](
	parser func(name string, value T, usage string) *T,
	name string,
	val T,
	desc string,
) InputFlag[T] {
	return InputFlag[T]{
		name:         name,
		defaultValue: val,
		description:  desc,
		value:        parser(name, val, desc),
	}
}

func (f *InputFlag[T]) Value() T {
	return *f.value
}

type TodoCliInput struct {
	addTodo      InputFlag[bool]
	showList     InputFlag[bool]
	completeTodo InputFlag[int]
}

func NewInput() *TodoCliInput {

	addTodo := NewInputFlag(
		flag.Bool,
		"add",
		false,
		"Add new todo from stdin (if provided) or args",
	)

	showList := NewInputFlag(
		flag.Bool,
		"list",
		false,
		"Show list of all todos",
	)

	completeTodo := NewInputFlag(
		flag.Int,
		"complete",
		-1,
		"Complete given todo by index",
	)

	return &TodoCliInput{
		addTodo:      addTodo,
		showList:     showList,
		completeTodo: completeTodo,
	}
}

func (t *TodoCliInput) Parse() {
	flag.Parse()
}
