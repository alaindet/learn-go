package main

import "fmt"

type FPerson struct {
	name, position string
}

type fpersonMod func(*FPerson)

type FPersonBuilder struct {
	actions []fpersonMod
}

// This is a sub-builer
func (b *FPersonBuilder) Named(name string) *FPersonBuilder {
	b.actions = append(b.actions, func(p *FPerson) {
		p.name = name
	})
	return b
}

// This is a sub-builer
func (b *FPersonBuilder) WorksAsA(position string) *FPersonBuilder {
	b.actions = append(b.actions, func(p *FPerson) {
		p.position = position
	})
	return b
}

// All sub-builders have been registered, not executed
// This function executes it all in order
func (b *FPersonBuilder) Build() *FPerson {
	person := &FPerson{}
	for _, action := range b.actions {
		action(person)
	}
	return person
}

func functionalBuilderExample() {
	b := FPersonBuilder{}
	p := b.Named("Alain").WorksAsA("Developer").Build()
	fmt.Println("person", *p)
}
