/*
The prototype pattern creates objects starting from an existing prototype which
is then customized

It's useful when the object is complex and/or there are lots of defaults

It's based on the principle of deep copy or cloning, which is copying an object
to create a new object perfectly identical to the first one, without any reference
to the original object
*/

package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func NewPerson(name string, address *Address, friends []string) *Person {
	return &Person{name, address, friends}
}

func (p *Person) DeepCopy() *Person {
	q := p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return q
}

// This DOES NOT WORK since *Address is a pointer
func wrongExample() {
	john := NewPerson(
		"John",
		&Address{"123 Foo Street", "London", "UK"},
		[]string{"Chris", "Matt"},
	)
	jane := john
	jane.Address.StreetAddress = "456 Bar Road"
	fmt.Println("\nwrongExample")
	fmt.Println(john.Address.StreetAddress) // 456 Bar Road
	fmt.Println(jane.Address.StreetAddress) // 456 Bar Road
}

// This works, but it does not scale
func workingExample() {
	john := NewPerson(
		"John",
		&Address{"123 Foo Street", "London", "UK"},
		[]string{"Chris", "Matt"},
	)
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}
	jane.Address.StreetAddress = "456 Bar Road"
	fmt.Println("\nworkingExample")
	fmt.Println(john.Address.StreetAddress) // 123 Foo Street
	fmt.Println(jane.Address.StreetAddress) // 456 Bar Road
}

func main() {
	wrongExample()
	workingExample()
	copyExample()
	serializationExample()
	prototypeFactoryExample()
}
