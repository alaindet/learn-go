package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type solid interface {
	volume() float64
}

// This interface is embedding the "shape" and "solid" interfaces!
type geometry interface {
	shape
	solid
	getColor() string
}

type cube struct {
	side  float64
	color string
}

func (c cube) area() float64 {
	return 6 * c.side * c.side
}

func (c cube) volume() float64 {
	return c.side * c.side * c.side
}

func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	area := g.area()
	volume := g.volume()
	return area, volume
}

func main() {
	c := cube{side: 3.0}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v) // Area: 54, Volume: 27
}
