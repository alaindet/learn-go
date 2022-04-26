package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		// Test case 1
		{
			Name:     "Struct with one string field",
			Input:    struct{ Name string }{"Chris"},
			Expected: []string{"Chris"},
		},
		// Test case 2
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			Expected: []string{"Chris"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			var result []string

			walk(testCase.Input, func(input string) {
				result = append(result, input)
			})

			if !reflect.DeepEqual(result, testCase.Expected) {
				t.Errorf("Result: %v, Expected: %v", result, testCase.Expected)
			}
		})
	}
}
