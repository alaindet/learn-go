package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{W: 10.0, H: 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("Result: %.2f Expected: %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {

	// This is table driven testing
	// https://github.com/golang/go/wiki/TableDrivenTests

	testCases := []struct {
		input    Shape
		expected float64
	}{
		{
			input:    Rectangle{W: 12.0, H: 6.0},
			expected: 72.0,
		},
		{
			input:    Circle{R: 10.0},
			expected: 314.1592653589793,
		},
		{
			input:    Triangle{B: 10.0, H: 5.0},
			expected: 25.0,
		},
	}

	for _, testCase := range testCases {
		result := testCase.input.Area()
		if result != testCase.expected {
			t.Errorf(
				"Input: %#v Result: %.2f Expected: %.2f",
				testCase.input,
				result,
				testCase.expected,
			)
		}
	}
}
