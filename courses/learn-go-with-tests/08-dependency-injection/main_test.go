package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	result := buffer.String()
	expected := "Hello, Chris"

	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}
