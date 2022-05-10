package main

import (
	"fmt"
	"time"
)

type DB struct {
	todos []Todo
}

func NewDB() *DB {
	return &DB{
		[]Todo{
			{
				Id:          "create",
				Description: "Create another Todo app",
				Done:        true,
				Due:         time.Now().Add(time.Hour * 48),
			},
			{
				Id:          "benchmark",
				Description: "Use it as a benchmark to try Go",
				Done:        false,
				Due:         time.Now().Add(time.Hour * 24),
			},
		},
	}
}

func (db *DB) GetTodos() []Todo {
	return db.todos
}

func (db *DB) GetTodoById(id string) (Todo, error) {
	for _, todo := range db.todos {
		if todo.Id == id {
			return todo, nil
		}
	}

	return Todo{}, fmt.Errorf("Todo with ID %q not found", id)
}
