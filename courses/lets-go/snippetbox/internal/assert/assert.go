package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	assertion := actual == expected
	if !assertion {
		t.Errorf("Value %v should be to be equal to %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, expected string) {
	t.Helper()
	assertion := strings.Contains(actual, expected)
	if !assertion {
		t.Errorf("Value %q should contain %q", actual, expected)
	}
}
