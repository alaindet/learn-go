package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	result := buffer.String()
	expected := "3\n2\n1\nGo!"

	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}

	if spySleeper.Calls != 4 {
		t.Errorf(
			"Not enough calls to sleeper, expected 4 got %d",
			spySleeper.Calls,
		)
	}
}
