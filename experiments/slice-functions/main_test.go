package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceMap(t *testing.T) {
	input := []int{1, 2, 3, 4}
	mapper := func(item int, i int) string {
		return fmt.Sprintf("%d%d", item, item)
	}
	result := sliceMap(input, mapper)
	expected := []string{"11", "22", "33", "44"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}

func TestSliceFindIndex(t *testing.T) {

	input := []int{10, 20, 30, 40}

	t.Run("should find item in array", func(t *testing.T) {
		finder := func(item int, i int) bool {
			return item == 30
		}
		result := sliceFindIndex(input, finder)
		expected := 2

		if result != expected {
			t.Errorf("Result: %v, Expected: %v", result, expected)
		}
	})

	t.Run("should not find item in array", func(t *testing.T) {
		finder := func(item int, i int) bool {
			return item == 42
		}
		result := sliceFindIndex(input, finder)
		expected := -1

		if result != expected {
			t.Errorf("Result: %v, Expected: %v", result, expected)
		}
	})
}

func TestSliceFind(t *testing.T) {

	input := []string{"10", "20", "300", "40", "5"}

	t.Run("should find item in array", func(t *testing.T) {
		finder := func(item string, i int) bool {
			return len(item) == 3
		}
		result, _ := sliceFind(input, finder)
		expected := "300"

		if result != expected {
			t.Errorf("Result: %v, Expected: %v", result, expected)
		}
	})

	t.Run("should not find item in array", func(t *testing.T) {
		finder := func(item string, i int) bool {
			return len(item) == 5
		}
		_, ok := sliceFind(input, finder)

		if ok == true {
			t.Errorf("Should not find the item, found instead")
		}
	})
}

func TestSliceFilter(t *testing.T) {
	input := []int{1, 2, 3, 4}
	isOdd := func(item int, i int) bool {
		return item%2 != 0
	}
	result := sliceFilter(input, isOdd)
	expected := []int{1, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}
