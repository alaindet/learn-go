package fibonacci

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Input    int
	Expected uint64
}

func TestFibonacci(t *testing.T) {
	testCases := []TestCase{
		{Input: 0, Expected: 0},
		{Input: 1, Expected: 1},
		{Input: 5, Expected: 5},
		{Input: 10, Expected: 55},
		{Input: 20, Expected: 6765},
		{Input: 60, Expected: 1548008755920},
		{Input: 80, Expected: 23416728348467685},
		{Input: 90, Expected: 2880067194370816120},
		{Input: 93, Expected: 12200160415121876738},
	}

	for _, testCase := range testCases {
		name := fmt.Sprintf("Fibonacci(%d)", testCase.Input)
		t.Run(name, func(t *testing.T) {
			t.Helper()
			result := Fibonacci(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("got %d expected %d", testCase.Input, testCase.Expected)
			}
		})
	}
}
