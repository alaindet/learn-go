package main

import (
	"math"
	"testing"
	"time"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(1970, time.January, 1, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

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
			result := toFixed(secondsInRadians(c.time), 5)
			expected := toFixed(c.angle, 5)

			if result != expected {
				t.Fatalf("got %v but wanted %v radians", expected, result)
			}
		})
	}
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1970, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	result := SecondHand(tm)
// 	expected := Point{X: Center, Y: Center + SecondHandLength}

// 	if result != expected {
// 		t.Errorf("Got %v expected %v", result, expected)
// 	}
// }
