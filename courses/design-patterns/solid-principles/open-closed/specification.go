package main

// The "specification" pattern shows the open-closed principle
type Specification interface {
	IsSatisfied(p *Product) bool
}
