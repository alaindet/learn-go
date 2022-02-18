/////////////////////////////////
// Structs in Go
// Go Playground: https://play.golang.org/p/AgeB0sjDUWQ
/////////////////////////////////

package main

import "fmt"

func examples1() {

	// creating a struct type
	type book struct {
		title  string //the fields of the book struct
		author string //each field must be unique inside a struct
		year   int
	}

	// combining different fields of the same type on the same line
	type book1 struct {
		title, author string
		year, pages   int
	}

	// declaring, initializing and assigning a new book value, all in one step
	lastBook := book{"The Divine Comedy", "Dante Aligheri", 1320} //this is a struct literal and order matters
	fmt.Println(lastBook)

	// Declaring a new book value by specifying field: value (order doesn't matter)
	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	_ = bestBook

	//if we create a new struct value by omitting some fields they will be zero-valued according to their type
	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook) // => main.book{title:"Just a random book", author:"", year:0}

	// retrieving the value of a struct field
	fmt.Println(lastBook.title) // => The Divine Comedy

	// selecting a field that doesn't exist raises an error
	// pages := lastBook.pages // error -> lastBook.pages undefined (type book has no field or method pages)

	// updating a field
	lastBook.author = "The Best"
	lastBook.year = 2020
	fmt.Printf("lastBook: %+v\n", lastBook) // => lastBook: {title:The Divine Comedy author:The Best year:2020}
	// + modifier with %v  printed out both the field names and their values

	// comparing struct values
	// two struct values are equal if their corresponding fields are equal.
	randomBook := book{title: "Random Title", author: "John Doe", year: 100}
	fmt.Println(randomBook == lastBook) // => false

	// = creates a copy of a struct
	myBook := randomBook
	myBook.year = 2020              // modifying only myBook
	fmt.Println(myBook, randomBook) // => {Random Title John Doe 2020} {Random Title John Doe 100}
}

func example2() {

	// an anonymous struct is a struct with no explicitly defined struct type alias.
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	// =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30

	//** ANONYMOUS FIELDS **//

	// fields type becomes fields name.
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // => main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

	fmt.Println(b1.string) // => 1984 by George Orwell

	// mixing anonymous with named fields:
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}
	fmt.Printf("%#v\n", e) // => main.Employee1{name:"John", salary:40000, bool:false}
	e.bool = true          // changing a field

	fmt.Println(strings.Repeat("#", 10))

	//** EMBEDDED STRUCTS **//
	// An embedded struct is just a struct that acts like a field inside another struct.

	// define a new struct type
	type Contact struct {
		email, address string
		phone          int
	}

	// define a struct type that contains another struct as a field
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	// declaring a value of type Employee
	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   042324234,
		},
	}

	fmt.Printf("%+v\n", john)
	// => {name:John Keller salary:3000 contactInfo:{email:jkeller@company.com address:Street 20, London phone:295619381404}}

	// accessing a field
	fmt.Printf("Employee's salary: %d\n", john.salary)

	// accessing a field from the embedded struct
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // => Employee's email:jkeller@company.com

	// updating a field
	john.contactInfo.email = "new_email@thecompany.com"
	fmt.Printf("Employee's new email address:%s\n", john.contactInfo.email)
	// =>  Employee's new email address:new_email@thecompany.com
}
