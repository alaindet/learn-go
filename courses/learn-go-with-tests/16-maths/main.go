package main

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
	return Point{}
}

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * math.Pi / 30
}
