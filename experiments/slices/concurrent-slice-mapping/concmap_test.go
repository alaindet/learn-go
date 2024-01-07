package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcMap(t *testing.T) {

	t.Run("works with small slice of integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

		output := ConcMap(input, double, -1)
		slices.Sort(input)
		slices.Sort(output)
		assert.Equal(t, output, expected)
	})

	t.Run("works with larger slices", func(t *testing.T) {
		size := 100_000
		input := make([]int, 0, size)
		expected := make([]int, 0, size)
		for i := 0; i < size; i++ {
			input = append(input, i+1)
			expected = append(expected, double(i+1))
		}

		output := ConcMap(input, double, -1)
		slices.Sort(input)
		slices.Sort(output)
		assert.Equal(t, output, expected)
	})
}

func double(n int) int {
	return n * 2
}
