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
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), 45 * math.Pi / 30},
		{simpleTime(0, 0, 7), 7 * math.Pi / 30},
	}

	for _, c := range cases {
		name := testName(c.time)
		t.Run(name, func(t *testing.T) {
			result := secondsInRadians(c.time)
			expected := c.angle

			if !circaEqualFloats(result, expected, 1e-7) {
				t.Fatalf("got %v but wanted %v radians", expected, result)
			}
		})
	}
}

func TestSecondHandVector(t *testing.T) {

	testCases := []struct {
		time     time.Time
		expected Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, testCase := range testCases {
		t.Run(testName(testCase.time), func(t *testing.T) {
			result := secondHandPoint(testCase.time)
			if !circaEqualPoints(result, testCase.expected) {
				t.Fatalf("Wanted %v Point, but got %v", testCase.expected, result)
			}
		})
	}
}
