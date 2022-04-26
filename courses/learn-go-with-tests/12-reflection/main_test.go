package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type PersonShortInfo struct {
	Name string
	Age  int
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		// Test case 1
		{
			Name:     "struct with one string field",
			Input:    struct{ Name string }{"Foo"},
			Expected: []string{"Foo"},
		},

		// Test case 2
		{
			Name:     "struct with non string field",
			Input:    PersonShortInfo{"Foo", 42},
			Expected: []string{"Foo"},
		},

		// Test case 3
		{
			Name:     "struct with nested fields",
			Input:    Person{"Foo", Profile{42, "London"}},
			Expected: []string{"Foo", "London"},
		},

		// Test case 4
		{
			Name:     "pointers to things",
			Input:    &Person{"Foo", Profile{42, "London"}},
			Expected: []string{"Foo", "London"},
		},

		// Test case 5
		{
			Name: "slices",
			Input: []Profile{
				{42, "London"},
				{69, "Roma"},
			},
			Expected: []string{"London", "Roma"},
		},

		// Test case 6
		{
			Name: "arrays",
			Input: [2]Profile{
				{42, "London"},
				{69, "Roma"},
			},
			Expected: []string{"London", "Roma"},
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
