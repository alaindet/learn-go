package main

func (a *AppContext) routes(routePrefix string) {

	route := func(path string) string {
		return routePrefix + "/" + path
	}

	// TODO: Grouping?
	a.fiber.Post(route("todos"), createTodoHandlerFn(a))
	a.fiber.Get(route("todos"), readTodosHandlerFn(a))
	a.fiber.Get(route("todos/:id"), readTodoHandlerFn(a))
	a.fiber.Put(route("todos/:id"), updateTodoHandlerFn(a))
	a.fiber.Delete(route("todos/:id"), deleteTodoHandlerFn(a))
}
