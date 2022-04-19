package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	check := func(t testing.TB, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: %q, Expected: %q", result, expected)
		}
	}

	t.Run(
		"Repeat 5 times",
		func(t *testing.T) {
			result := Repeat("a", 5)
			expected := "aaaaa"
			check(t, result, expected)
		},
	)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("foo", 3)
	fmt.Println(result)
	// Output: foofoofoo
}
