package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func createTodoHandler(c *fiber.Ctx) error {
	return c.SendString("Create todo")
}

func readTodosHandler(c *fiber.Ctx) error {
	return c.SendString("Read todos")
}

func readTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	message := fmt.Sprintf("Read todo: ID %s", id)
	return c.SendString(message)
}

func updateTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	message := fmt.Sprintf("Update todo: ID %s", id)
	return c.SendString(message)
}

func deleteTodoHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	message := fmt.Sprintf("Delete todo: ID %s", id)
	return c.SendString(message)
}
