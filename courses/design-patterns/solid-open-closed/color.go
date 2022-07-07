package main

type Color int

const (
	red Color = iota
	green
	blue
)

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}
