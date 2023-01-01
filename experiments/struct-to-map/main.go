package main

import (
	"fmt"
)

func main() {
	// Convert an anonymous struct to a map, recursively
	fmt.Printf("%+v\n", structToMap(struct {
		Name    string
		Hobbies []string
		Partner struct {
			Name    string
			Hobbies []string
		}
		Pets []struct {
			Name string
		}
	}{
		Name:    "Alice",
		Hobbies: []string{"Programming", "Hiking"},
		Partner: struct {
			Name    string
			Hobbies []string
		}{
			Name:    "Bernadette",
			Hobbies: []string{"Reading", "Traveling"},
		},
		Pets: []struct {
			Name string
		}{
			{Name: "Chilly"},
			{Name: "Dexter"},
		},
	}))
}
