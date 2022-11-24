/*
The factory design pattern is best used when creation logic of something is too
difficult, for example too many struct fields, too many dependencies

It simplifies creation of a lot of similar objects as compared to the builder
pattern, which specializes into the creation of one object at a time

The factory function (the constructor) helps creating structs and can perform
additional creation logic, like validation

If a factory returns an interface, data is actually incapsulated
*/

package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello, I am %s.\n", p.name)
}

type oldPerson struct {
	name string
	age  int
}

func (p *oldPerson) SayHello() {
	fmt.Printf("Hello, I am %s and I'm old!\n", p.name)
}

// Notice we're returning an interface
func NewPerson(name string, age int) Person {
	if age > 70 {
		return &oldPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("John", 20)
	p.SayHello()

	old := NewPerson("Wilbur", 85)
	old.SayHello()

	factoryGeneratorExample()
	prototypeFactoryExample()
}
