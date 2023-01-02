package main

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {

	data := struct {
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
	}

	result := StructToMap(data)

	// TODO
	fmt.Printf("%+v", result)

	// assertEqual(t, result["Name"].(string), "Alice")
	// assertEqual(t, len(result["Hobbies"].([]string)), 2)
}
