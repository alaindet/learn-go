/*
Bridge Design Pattern

Connects two components together via an abstraction (the bridge), so that the
communication between components is simplified. It avoids a specific problem
called "cartesian product complexity explosion"
*/

package main

import "fmt"

// Shapes: Circle, Square
// Image type: Raster, Vector

// This is the bridge interface!
type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(side float32)
}

type VectorRenderer struct {
	// ...
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a vector circle with radius", radius)
}

func (v *VectorRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing a vector square with side", side)
}

type RasterRenderer struct {
	// ...
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a raster circle with radius", radius)
}

func (r *RasterRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing a raster square with side", side)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Square struct {
	renderer Renderer
	side     float32
}

func NewSquare(renderer Renderer, side float32) *Square {
	return &Square{renderer, side}
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

func (s *Square) Resize(factor float32) {
	s.side *= factor
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	c1 := NewCircle(&raster, 5)
	c1.Draw()
	c1.Resize(2)
	c1.Draw()

	s1 := NewSquare(&raster, 5)
	s1.Draw()
	s1.Resize(1.5)
	s1.Draw()

	c2 := NewCircle(&vector, 10)
	c2.Draw()
	c2.Resize(2)
	c2.Draw()

	s2 := NewSquare(&vector, 10)
	s2.Draw()
	s2.Resize(1.5)
	s2.Draw()
}
