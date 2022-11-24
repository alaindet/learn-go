package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Functional approach
// This approach helps immutability and predictability as factory parameters
// cannot be changed after factory creation
// This is easier for exposing APIs
func NewFunctionalEmployeeFactory(pos string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, pos, income}
	}
}

func functionalApproachExample() {
	developerFactory := NewFunctionalEmployeeFactory("developer", 20000)
	managerFactory := NewFunctionalEmployeeFactory("manager", 40000)

	dev1 := developerFactory("Adam")
	dev2 := developerFactory("Bernie")
	mang1 := managerFactory("Chandler")

	fmt.Println(dev1, dev2, mang1)
}

// Structural approach
// This approach allows you to change position and/or income down
// after creating the factory
// This is harder for exposing APIs since users must know how the specialized
// factory works
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewStructuralEmployeeFactory(pos string, income int) *EmployeeFactory {
	return &EmployeeFactory{pos, income}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func structualApproachExample() {
	developerFactory := NewStructuralEmployeeFactory("developer", 20000)
	managerFactory := NewStructuralEmployeeFactory("manager", 40000)

	dev1 := developerFactory.Create("Adam")
	dev2 := developerFactory.Create("Bernie")
	mang1 := managerFactory.Create("Chandler")

	fmt.Println(dev1, dev2, mang1)
}

func factoryGeneratorExample() {
	functionalApproachExample()
	structualApproachExample()
}
