package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	check := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	}

	t.Run(
		"collection of any size",
		func(t *testing.T) {
			arg := []int{1, 2, 3}
			result := Sum(arg)
			expected := 6
			check(t, result, expected)
		},
	)
}

func TestSumAll(t *testing.T) {

	check := func(t *testing.T, result, expected []int) {
		t.Helper()
		if reflect.DeepEqual(result, expected) {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	}

	t.Run(
		"sum a slice []int, each with a total",
		func(t *testing.T) {
			result := SumAll([]int{1, 2}, []int{0, 9})
			expected := []int{3, 9}
			check(t, result, expected)
		},
	)
}

func TestSumAllTails(t *testing.T) {

	check := func(t *testing.T, result, expected []int) {
		t.Helper()
		if reflect.DeepEqual(result, expected) {
			t.Errorf("Result: %d, Expected: %d", result, expected)
		}
	}

	t.Run(
		"make the sums of some slices",
		func(t *testing.T) {
			result := SumAllTails([]int{1, 2}, []int{0, 9})
			expected := []int{2, 9}
			check(t, result, expected)
		},
	)

	t.Run(
		"safely sum empty slices",
		func(t *testing.T) {
			result := SumAllTails([]int{}, []int{3, 4, 5})
			expected := []int{0, 9}
			check(t, result, expected)
		},
	)
}
