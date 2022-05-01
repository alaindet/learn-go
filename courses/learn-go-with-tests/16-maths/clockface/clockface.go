package clockface

import (
	"math"
	"time"
)

const (
	Center           = 150
	SecondHandLength = 90
	MinuteHandLength = 80
	HourHandLength   = 50
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * SecondHandLength, p.Y * SecondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip on Y axis
	p = Point{p.X + Center, p.Y + Center}                     // translate
	return p
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * math.Pi / 30

	// Equivalent, but should preserve floating-point precision (?)
	return math.Pi / (30 / float64(t.Second()))
}

// This assumes a "unit" circle of radius 1 centered in (0,0)
// Returns the coordinates of the second hand tip on the unit circle
func secondHandPoint(t time.Time) Point {
	rad := secondsInRadians(t)
	x := math.Sin(rad)
	y := math.Cos(rad)
	return Point{x, y}
}
