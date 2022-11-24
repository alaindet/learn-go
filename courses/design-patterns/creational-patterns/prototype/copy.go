package main

import "fmt"

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func printPerson(p *Person) {
	fmt.Printf(
		"Name: %s, Address: %+v, Friends: %+v\n",
		p.Name,
		p.Address,
		p.Friends,
	)
}

func copyExample() {
	john := NewPerson(
		"John",
		&Address{"123 Foo Street", "London", "UK"},
		[]string{"Chris", "Matt"},
	)

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "456 Bar Road"
	jane.Friends = append(jane.Friends, "Anne")

	fmt.Println("\ncopyExample")
	printPerson(john)
	printPerson(jane)
}
