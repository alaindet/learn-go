package main

import (
	"testing"
	"time"
)

func TestFriendlyDate(t *testing.T) {
	input := time.Date(2001, 2, 3, 4, 5, 6, 7, time.UTC)
	result := friendlyDate(input)
	expected := "03 Feb 2001 at 04:05"

	if result != expected {
		t.Errorf("Result: %q; Expected: %q", result, expected)
	}
}
