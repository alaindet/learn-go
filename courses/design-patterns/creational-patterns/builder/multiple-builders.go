package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

// PersonBuilder --------------------------------------------------------------
type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

// PersonAddressBuilder -------------------------------------------------------
type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

// PersonJobBuilder -----------------------------------------------------------
type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

// Example --------------------------------------------------------------------
func multipleBuildersExample() {

	// A complex builder using builders
	person := NewPersonBuilder(). // Init main PersonBuilder
					Lives().                 // Init inner PersonAddressBuilder
					At("123 Foo Street").    // Populate PersonAddressBuilder
					In("Some City").         // Populate PersonAddressBuilder
					WithPostcode("AA55BB").  // Populate PersonAddressBuilder
					Works().                 // Init inner PersonJobBuilder
					At("Some Company Inc."). // Populate PersonJobBuilder
					AsA("Programmer").       // Populate PersonJobBuilder
					Earning(20000).          // Populate PersonJobBuilder
					Build()                  // Back to PersonBuilder to get the final Person

	fmt.Printf("Person: %+v\n", person)

	// Equivalent
	pb := NewPersonBuilder()

	pab := pb.Lives()
	pab.At("123 Foo Street")
	pab.In("Some City")
	pab.WithPostcode("AA55BB")

	pjb := pb.Works()
	pjb.At("Some Company Inc.")
	pjb.AsA("Programmer")
	pjb.Earning(20000)

	person2 := pb.Build()
	fmt.Printf("Person: %+v\n", person2)
}
