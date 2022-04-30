package main

import (
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	input := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	expected := Point{X: 150, Y: 150 - 90}
	result := SecondHand(input)

	if result != expected {
		t.Errorf("result: %v expected: %v", result, expected)
	}
}
