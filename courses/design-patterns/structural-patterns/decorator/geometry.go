package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Rendering a circle with radius %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Square() string {
	return fmt.Sprintf("Rendering a square with side %.2f", s.Side)
}

// This is a decorator
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s and with color %s", c.Shape.Render(), c.Color)
}

// This is a decorator
type TransparentShape struct {
	Shape   Shape
	Opacity float32
}

func (t *TransparentShape) Render() string {
	opacity := t.Opacity * 100.0
	return fmt.Sprintf("%s and with opacity %.2f", t.Shape.Render(), opacity)
}

func geometryExample() {
	c := Circle{2}
	fmt.Println(c.Render())

	rc := ColoredShape{&c, "red"}
	fmt.Println(rc.Render())
	// rc.Resize() // <-- Cannot do it!

	tc := TransparentShape{&c, 0.66}
	fmt.Println(tc.Render())
}
