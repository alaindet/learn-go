package clockface

import (
	"math"
	"testing"
	"time"
)

func circaEqualFloats(a, b, precision float64) bool {
	return math.Abs(b-a) < precision
}

func circaEqualPoints(a, b Point) bool {
	return circaEqualFloats(a.X, b.X, 1e-7) && circaEqualFloats(a.Y, b.Y, 1e-7)
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(1970, time.January, 1, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

// This is a unit test
func TestSecondsInRadians(t *testing.T) {

	cases := []struct {
		input    time.Time
		expected float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), 45 * math.Pi / 30},
		{simpleTime(0, 0, 7), 7 * math.Pi / 30},
	}

	for _, c := range cases {
		name := testName(c.input)
		t.Run(name, func(t *testing.T) {
			result := secondsInRadians(c.input)
			expected := c.expected

			if !circaEqualFloats(result, expected, 1e-7) {
				t.Fatalf("got %v but wanted %v radians", expected, result)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {

	testCases := []struct {
		input    time.Time
		expected Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, testCase := range testCases {
		t.Run(testName(testCase.input), func(t *testing.T) {
			result := secondHandPoint(testCase.input)
			if !circaEqualPoints(result, testCase.expected) {
				t.Fatalf("Wanted %v Point, but got %v", testCase.expected, result)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		input    time.Time
		expected float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.input), func(t *testing.T) {
			result := minutesInRadians(c.input)
			if result != c.expected {
				t.Fatalf("Wanted %v radians, but got %v", c.expected, result)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		input    time.Time
		expected Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.input), func(t *testing.T) {
			result := minuteHandPoint(c.input)
			if !circaEqualPoints(result, c.expected) {
				t.Fatalf("Wanted %v Point, but got %v", c.expected, result)
			}
		})
	}
}
