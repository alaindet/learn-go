/*
Composite Design Pattern

The composite design pattern allows both simple objects and complex objects
to be treated uniformly as simple objects, where "complex objects" are composed
by aggregations of simple objects or inner complex objects

Ex.:
A car is both a component on a traffic'd road and an aggregation of components
(engine, brakes etc.). An engine is both a component for a car and an aggregation
of sub-components
*/
package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) print(b *strings.Builder, depth int) {
	b.WriteString(strings.Repeat("> ", depth))
	if len(g.Color) > 0 {
		b.WriteString(g.Color)
		b.WriteRune(' ')
	}
	b.WriteString(g.Name)
	b.WriteRune('\n')

	for _, child := range g.Children {
		child.print(b, depth+1)
	}
}

func (g *GraphicObject) String() string {
	b := strings.Builder{}
	g.print(&b, 0)
	return b.String()
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func drawingGroupsExample() {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Blue"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(drawing.Children, *NewCircle("Green"))
	group.Children = append(drawing.Children, *NewCircle("Light Blue"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
	/*
		My Drawing
		> Red Circle
		> Blue Square
		> Group 1
		> > Red Circle
		> > Blue Square
		> > Light Blue Circle
	*/
}

func main() {
	drawingGroupsExample()
	neuralNetworkExample()
}
