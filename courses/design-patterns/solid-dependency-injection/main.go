package main

import "fmt"

/*
Dependency Injection Principle
High-level modules should not depend on low-level modules and viceversa
Instead, they should both depend on abstractions
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type PersonConnection struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-level module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// This is a low level module as it has storage
type Relationships struct {
	relationships []PersonConnection
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	childToParent := PersonConnection{child, Child, parent}
	parentToChild := PersonConnection{parent, Parent, child}
	r.relationships = append(r.relationships, childToParent, parentToChild)
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0, len(r.relationships))

	for _, rel := range r.relationships {
		if rel.from.name == name && rel.relationship == Parent {
			result = append(result, &Person{rel.to.name})
		}
	}

	return result
}

// This is a high level module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate(
	ofName string,
	ofRelationship Relationship,
) []*Person {
	switch ofRelationship {
	case Parent:
		return r.browser.FindAllChildrenOf(ofName)
	case Child:
		return make([]*Person, 0) // TODO
	case Sibling:
		return make([]*Person, 0) // TODO
	default:
		return make([]*Person, 0)
	}
}

func main() {
	parent := &Person{"John"}
	child1 := &Person{"Mary"}
	child2 := &Person{"Gerald"}

	rels := Relationships{}
	rels.AddParentAndChild(parent, child1)
	rels.AddParentAndChild(parent, child2)

	r := Research{&rels}
	parentName := "John"
	found := r.Investigate(parentName, Parent)

	for _, p := range found {
		fmt.Printf("%s is child of %s\n", p.name, parentName)
	}
}
