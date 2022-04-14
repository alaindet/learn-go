package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func updateTodoHandlerFn(a *AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		id, err := strconv.Atoi(c.Params("id"))
		_ = err

		todo, err := updateTodo(a.db, &UpdateTodoDto{
			Id:     int64(id),
			Name:   "Updated name", // TODO: Get from request
			IsDone: true,           // TODO: Get from request
		})

		// TODO: Check HTTP error
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				fmt.Sprintf("Cannot update todo with id %d", id),
			)
		}

		return c.SendString(
			fmt.Sprintf("Todo %q updated", todo.Name),
		)
	}
}
