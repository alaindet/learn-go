package main

import (
	"fmt"
	"math"
)

// This is an interface
type shape interface {
	area() float64
	perimeter() float64
}

// Circle ---------------------------------------------------------------------
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	width, height float64
}

// Rectangle ------------------------------------------------------------------
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

/**
 * This is "programming against interfaces" in action
 */
func print(s shape) {
	fmt.Printf(
		"Shape: %T, Area: %.2f, Perimeter: % .2f\n",
		s,
		s.area(),
		s.perimeter(),
	)
}

func main() {
	c := circle{radius: 10.0}
	r := rectangle{width: 16.0, height: 9.0}

	print(c) // Shape: main.circle, Area: 314.16, Perimeter:  62.83
	print(r) // Shape: main.rectangle, Area: 144.00, Perimeter:  50.00
}
