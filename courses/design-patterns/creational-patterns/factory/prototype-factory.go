package main

import "fmt"

/*
A prototype is like a template to be reused by the factory
*/

const (
	Developer = iota
	Manager
)

// // from factory-generators.go
// type Employee struct {
// 	Name, Position string
// 	AnnualIncome   int
// }

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 20000}
	case Manager:
		return &Employee{"", "manager", 30000}
	default:
		panic("Unsupported role")
	}
}

func prototypeFactoryExample() {
	m := NewEmployee(Manager)
	m.Name = "Geoffrey"
	fmt.Println("manager", m)

	d := NewEmployee(Developer)
	d.Name = "John"
	fmt.Println("developer", d)
}
