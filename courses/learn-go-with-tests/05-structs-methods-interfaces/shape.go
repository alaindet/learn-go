package main

import "math"

type Rectangle struct {
	W float64
	H float64
}

type Circle struct {
	R float64
}

type Triangle struct {
	B float64
	H float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (c Circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (t Triangle) Area() float64 {
	return t.B * t.H * 0.5
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.W + r.H)
}
