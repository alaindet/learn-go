package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func readTodoHandlerFn(a *AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		id, err := strconv.Atoi(c.Params("id"))
		_ = err

		todo, err := getTodoById(a.db, int64(id))

		// TODO: Check HTTP error
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				fmt.Sprintf("Cannot find todo with id %d", id),
			)
		}

		// TODO: Send todo data as JSON
		return c.SendString(
			fmt.Sprintf("Read todo: Name %s", todo.Name),
		)
	}
}

func readTodosHandlerFn(a *AppContext) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		todos, err := getTodos(a.db)

		// TODO: Check HTTP error
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Cannot find todos",
			)
		}

		// TODO: Send todos data as JSON
		return c.SendString(
			fmt.Sprintf("Read todos (%d)", len(todos)),
		)
	}
}
