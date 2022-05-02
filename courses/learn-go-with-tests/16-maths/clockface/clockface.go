package clockface

import (
	"fmt"
	"io"
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
	return math.Pi / (30 / float64(t.Second()))
	// ^^ Equivalent, but should preserve floating-point precision (?)
}

// This assumes a "unit" circle of radius 1 centered in (0,0)
// Returns the coordinates of the second hand tip on the unit circle
func secondHandPoint(t time.Time) Point {
	rad := secondsInRadians(t)
	x := math.Sin(rad)
	y := math.Cos(rad)
	return Point{x, y}
}

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p = Point{p.X * SecondHandLength, p.Y * SecondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + Center, p.Y + Center}                     // translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// This writes the image of a clock in SVG showing time t in writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
