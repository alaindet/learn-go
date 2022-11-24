package main

type Size int

const (
	small Size = iota
	medium
	large
)

type SizeSpecification struct {
	Size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == c.Size
}
