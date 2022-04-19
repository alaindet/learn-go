package main

import "testing"

func TestSum(t *testing.T) {

	check := func(t testing.TB, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: %q, Expected: %q", result, expected)
		}
	}

	numbers := [5]int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15
	check(t, result, expected)
}
