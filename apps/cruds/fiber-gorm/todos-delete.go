package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func deleteTodoHandlerFn(a *AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		id, err := strconv.Atoi(c.Params("id"))
		_ = err

		todo, err := deleteTodo(a.db, &DeleteTodoDto{
			Id: int64(id),
		})

		// TODO: Check HTTP error
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				fmt.Sprintf("Cannot delete todo with id %d", id),
			)
		}

		return c.SendString(
			fmt.Sprintf("Todo %q deleted", todo.Name),
		)
	}
}
