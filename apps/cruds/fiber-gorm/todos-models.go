package main

type Todo struct {
	Id     int64
	Name   string
	IsDone bool
}

type CreateTodoDto struct {
	Name string
}

type DeleteTodoDto struct {
	Id int64
}

type UpdateTodoDto struct {
	Id     int64
	Name   string
	IsDone bool
}
