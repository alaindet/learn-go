/*
A prototype factory allows for minor customizations but most of the object
is left intact
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Employee struct {
	Name   string
	Office Address
}

func newEmployee(proto *Employee, name string) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	// Other customizations here...
	return result
}

func (e *Employee) DeepCopy() *Employee {
	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)
	_ = enc.Encode(e)

	dec := gob.NewDecoder(buffer)
	newEmployee := &Employee{}
	_ = dec.Decode(newEmployee)

	return newEmployee
}

var mainOffice = Employee{
	"REPLACE_ME",
	Address{"123 Foo Road", "London", "UK"},
}

var auxOffice = Employee{
	"REPLACE_ME",
	Address{"456 Bar Road", "London", "UK"},
}

func NewMainOfficeEmployee(name string) *Employee {
	return newEmployee(&mainOffice, name)
}

func NewAuxiliaryOfficeEmployee(name string) *Employee {
	return newEmployee(&mainOffice, name)
}

func prototypeFactoryExample() {
	john := NewMainOfficeEmployee("John")
	jane := NewMainOfficeEmployee("Jane")
	margaret := NewAuxiliaryOfficeEmployee("Margaret")
	william := NewAuxiliaryOfficeEmployee("William")

	fmt.Println("prototypeFactoryExample")
	fmt.Println("John", john)
	fmt.Println("Jane", jane)
	fmt.Println("Margaret", margaret)
	fmt.Println("William", william)
}
