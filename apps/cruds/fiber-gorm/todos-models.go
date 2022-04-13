package main

type Todo struct {
	Id     int64
	Name   string
	IsDone bool
}

type CreateTodoDto struct {
	Name string
}
