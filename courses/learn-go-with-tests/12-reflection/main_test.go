package main

import "testing"

func TestWalk(t *testing.T) {

	expected := "Foo"
	var result []string

	x := struct {
		Name string
	}{Name: expected}

	walk(x, func(input string) {
		result = append(result, input)
	})

	if len(result) != 1 {
		t.Errorf(
			"wrong number of function calls, got %d want %d",
			len(result),
			1,
		)
	}
}
