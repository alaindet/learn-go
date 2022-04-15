package main

import "fmt"

func getTodoById(id string) (*Todo, error) {

	if id == "69" {
		return nil, fmt.Errorf("42")
	}

	todo := &Todo{
		Id:     "1",
		Name:   "A new thing to do",
		IsDone: false,
	}

	return todo, nil
}
