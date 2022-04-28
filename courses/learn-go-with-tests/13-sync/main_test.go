package main

import "testing"

func assertCounter(t testing.TB, c Counter, expected int) {
	t.Helper()
	result := c.Value()
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

func TestCounter(t *testing.T) {
	t.Run("basic synchronous counter", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()
		assertCounter(t, counter, 3)
	})
}
