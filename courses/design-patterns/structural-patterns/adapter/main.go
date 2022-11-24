/*
The Adapter Design Pattern transforms objects with an interface A to
an interface B via a software layer (the adapter), in order to conform to some
specification

Adapters should be immutable, i.e. memoizable

Examples
- Merging data from two different APIs
- Adapting an existing third-party library to an existing codebase
*/

package main

import "fmt"

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1  // Make it 0-based
	height -= 1 // Make it 0-based

	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{width, 0, width, height},
			{width, height, 0, height},
			{0, height, 0, 0},
		},
	}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

// This is the adapter
type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			v.points = append(v.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			v.points = append(v.points, Point{x, top})
		}
	}

	fmt.Printf("Generated %d points\n", len(v.points))
}

func VectorToRaster(v *VectorImage) RasterImage {
	adapter := &vectorToRasterAdapter{}

	for _, line := range v.Lines {
		adapter.addLine(line)
	}

	return adapter
}

func main() {
	rect := NewRectangle(6, 4)
	raster := VectorToRaster(rect)
	fmt.Print(DrawPoints(raster))
}
