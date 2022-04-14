package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func createTodoHandlerFn(a *AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		todo, err := a.repositories.todos.create(&CreateTodoDto{
			Name: "A new todo item", // TODO: Get it from the body
		})

		// TODO: Check HTTP error
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				fmt.Sprintf("Cannot create todo with name %q", todo.Name),
			)
		}

		return c.SendString(
			fmt.Sprintf("New todo created with ID %d", todo.Id),
		)
	}
}
