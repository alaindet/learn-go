package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, c *Counter, expected int) {
	t.Helper()
	result := c.Value()
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

func TestCounter(t *testing.T) {
	t.Run("basic synchronous counter", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()
		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expected := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expected)

		for i := 0; i < expected; i++ {
			go func(wg *sync.WaitGroup, c *Counter) {
				c.Increment()
				wg.Done()
			}(&wg, counter)
		}
		wg.Wait()

		assertCounter(t, counter, expected)
	})
}
