package main

import "flag"

type InputFlag[T any] struct {
	name         string
	defaultValue T
	description  string
	value        T
	parser       func(name string, value T, description string) *T
}

type TodoCliInput struct {
	addTodo      InputFlag[string]
	showList     InputFlag[bool]
	completeTodo InputFlag[int]
}

func NewInput() *TodoCliInput {

	addTodo := InputFlag[string]{
		name:         "todo",
		defaultValue: "",
		description:  "New todo to be included in the To Do list",
		parser:       flag.String,
	}

	showList := InputFlag[bool]{
		name:         "list",
		defaultValue: false,
		description:  "Show list of all todos",
		parser:       flag.Bool,
	}

	completeTodo := InputFlag[int]{
		name:         "complete",
		defaultValue: -1,
		description:  "Complete given todo by index",
		parser:       flag.Int,
	}

	return &TodoCliInput{
		addTodo:      addTodo,
		showList:     showList,
		completeTodo: completeTodo,
	}
}

func (t *TodoCliInput) Parse() {

	at := t.addTodo
	t.addTodo.value = *at.parser(at.name, at.defaultValue, at.description)

	sl := t.showList
	t.showList.value = *sl.parser(sl.name, sl.defaultValue, sl.description)

	ct := t.completeTodo
	t.completeTodo.value = *ct.parser(ct.name, ct.defaultValue, ct.description)

	flag.Parse()
}
