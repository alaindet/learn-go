package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected %v, but got %v", expected, result)
	}
}

func ExampleAdd() {
	// The comment at the bottom is semantic for the example!
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}
