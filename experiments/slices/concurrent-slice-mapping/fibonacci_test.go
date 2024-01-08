package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{10, 55},
		{15, 610},
		{18, 2584},
		{19, 4181},
		// From https://r-knott.surrey.ac.uk/Fibonacci/fibtable.html
		{90, 2880067194370816120},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("Fibonacci term #%d", testCase.input)
		t.Run(testName, func(t *testing.T) {
			outcome := Fibonacci(testCase.input)
			assert.Equal(t, testCase.expected, outcome)
		})
	}
}
