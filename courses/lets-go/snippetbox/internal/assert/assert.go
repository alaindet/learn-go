package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {

	// Marks this as an helper that can be skipped on logs
	t.Helper()

	assertion := actual == expected

	if !assertion {
		t.Errorf("Result: %v; Expected:: %v", actual, expected)
	}
}
