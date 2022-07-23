package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(side int) *Square {
	return &Square{Rectangle{side, side}}
}

// This breaks the Liskov Substitution Principle since UseSized(), which uses the
// Sized interface, doesn't work with Square because of this implementation
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// This breaks the Liskov Substitution Principle since UseSized(), which uses the
// Sized interface, doesn't work with Square because of this implementation
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseSized(s Sized) {
	w := s.GetWidth()
	s.SetHeight(10)
	expected := w * 10
	result := s.GetWidth() * s.GetHeight()

	if result != expected {
		fmt.Printf("ERROR: Expected %d, got %d\n", expected, result)
	} else {
		fmt.Println("PASSED")
	}
}

func main() {
	r := &Rectangle{2, 3}
	UseSized(r) // PASSED

	s := NewSquare(5)
	UseSized(s) // ERROR: Expected 50, got 100
}
