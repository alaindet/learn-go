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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
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

	// Test for maps
	t.Run("with maps", func(t *testing.T) {
		input := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(input, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	// Test for channels
	t.Run("channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func(outerCh chan Profile) {
			outerCh <- Profile{42, "London"}
			outerCh <- Profile{69, "Roma"}
			close(outerCh)
		}(ch)

		var result []string
		walk(ch, func(input string) {
			result = append(result, input)
		})

		expected := []string{"London", "Roma"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})

	// Test for functions
	t.Run("function", func(t *testing.T) {

		f := func() (Profile, Profile) {
			return Profile{42, "London"}, Profile{69, "Roma"}
		}

		var result []string
		walk(f, func(input string) {
			result = append(result, input)
		})

		expected := []string{"London", "Roma"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})
}
