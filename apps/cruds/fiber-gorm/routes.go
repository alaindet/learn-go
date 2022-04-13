package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var routePrefix string = ""

func setRoutePrefix(c config) {
	routePrefix = fmt.Sprintf("/api/v%s", c.VERSION)
}

func route(urlPattern string) string {
	return fmt.Sprintf("%s/%s", routePrefix, urlPattern)
}

func addStaticServing(app *fiber.App) *fiber.App {
	app.Static("/assets", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     false, // No file streaming
		Browse:        false, // No directory browsing
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	return app
}

func setupRoutes(app *fiber.App, c config) *fiber.App {

	setRoutePrefix(c)
	addStaticServing(app)

	// TODO: Grouping?
	app.Post(route("todos"), createTodoHandler)
	app.Get(route("todos"), readTodosHandler)
	app.Get(route("todos/:id"), readTodoHandler)
	app.Put(route("todos/:id"), updateTodoHandler)
	app.Delete(route("todos/:id"), deleteTodoHandler)

	// Add other routes here...

	return app
}
