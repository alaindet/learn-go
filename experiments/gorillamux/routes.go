package main

var routes = []Route{
	{
		Name:    "Index",
		Method:  "GET",
		Path:    "/",
		Handler: IndexHandler,
	},
	{
		Name:    "GetTodos",
		Method:  "GET",
		Path:    "/todos",
		Handler: GetTodosHandler,
	},
	{
		Name:    "GetTodo",
		Method:  "GET",
		Path:    "/todos/{todoid}",
		Handler: GetTodoHandler,
	},
}
