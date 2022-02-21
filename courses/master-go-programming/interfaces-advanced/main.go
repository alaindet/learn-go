package main

import (
	"fmt"
	"log"
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

func (c circle) volume() float64 {
	return (4 / 3) * math.Pi * c.radius * c.radius * c.radius
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

func lessonDynamicType() {
	var s shape
	fmt.Printf("%T\n", s) // <nil>

	ball := circle{radius: 10.0}
	s = ball // Shape: main.circle, Area: 314.16, Perimeter:  62.83

	print(s)
	fmt.Printf("%T\n", s) // main.circle

	frame := rectangle{width: 16.0, height: 9.0} // This is a concrete value
	s = frame
	fmt.Printf("%T\n", s) // main.rectangle
}

func lessonTypeAssertions() {
	var s shape = circle{radius: 24.5}
	fmt.Printf("%T\n", s) // main.circle
	// s.volume()            // Error: type shape has no field or method volume

	// Type assertion (Checks at runtime if shape can be converted to circle)
	c, ok := s.(circle)

	if !ok {
		log.Fatal("Error")
	}

	volume := c.volume()
	fmt.Printf("Volume: %.2f\n", volume) // Volume: 46200.65
}

func lessonTypeSwitches() {
	// s is a circle now
	var s shape = circle{radius: 24.5}

	// s became a rectangle
	s = rectangle{width: 16., height: 9.}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v is a circle\n", value)
	case rectangle:
		fmt.Printf("%#v is a rectangle\n", value)
	}
	// main.rectangle{width:16, height:9} is a rectangle
}

/**
 * Any type is an interface
 * It can represent any value
 * GO interface{} ~= TypeScript any
 * func hello(a ...interface{}) means "Any number of args of any type"
 */
type theEmptyInterface interface{}

type person struct {
	info interface{}
}

func lessonEmptyInterface() {
	var empty theEmptyInterface

	// "empty" variable can hold any value
	empty = 5
	empty = "hello world"
	empty = []string{"Hello", "World"}

	// But it cannot be used on len() since no type is guaranteed
	fmt.Printf("%T\n", empty) // []string
	// fmt.Println(len(empty))   // Error: Invalid argument for len
	emptyAsSlice := empty.([]string)
	fmt.Println(len(emptyAsSlice)) // 2

	// type "person" has "info" field in which you can store anything!
	aGuy := person{}
	aGuy.info = "I am Iron Man"
	aGuy.info = 42
	aGuy.info = []string{"I", "am", "Iron", "Man"}
	aGuy.info = aGuy // Store itself!
	// ...
}

func main() {
	lessonDynamicType()
	lessonTypeAssertions()
	lessonTypeSwitches()
	lessonEmptyInterface()
}
