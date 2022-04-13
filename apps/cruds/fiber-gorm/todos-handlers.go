package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func createTodo(db *sql.DB, dto *CreateTodoDto) (*Todo, error) {
	result, err := db.Exec(
		`INSERT INTO todos (name) VALUES (?)`,
		dto.Name,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	todo := &Todo{
		Id:     id,
		Name:   dto.Name,
		IsDone: false,
	}

	return todo, nil
}

func createTodoHandler(c *fiber.Ctx) error {
	db, err := connectToDatabase(c)

	// TODO: Send HTTP error
	if err != nil {
		return c.SendString("ERROR: Cannot create todo")
	}

	defer db.Close()

	todo, err := createTodo(
		db,
		&CreateTodoDto{
			Name: "A new todo item",
		},
	)

	// TODO: Send HTTP error
	if err != nil {
		return c.SendString("ERROR: Cannot create todo")
	}

	message := fmt.Sprintf("New todo created with ID %d", todo.Id)
	return c.SendString(message)
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
