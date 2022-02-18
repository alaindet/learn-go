package main

import (
	"fmt"
)

type book struct {
	title  string
	author string
	year   int
}

func createStructs() {
	// // These are just unrelated variables
	// book1Title := "La Divina Commedia"
	// book1Author := "Dante Alighieri"
	// book1Year := 1310

	// book2Title := "Macbeth"
	// book2Author := "William Shakespeare"
	// book2Year := 1606

	// NOT RECOMMENDED WAY - Here, order matters!
	bookBadPractice := book{
		"La Divina Commedia",
		"Dante Alighieri",
		1310,
	}
	_ = bookBadPractice

	// RECOMMENDED WAY - Explicit key/value assignments
	bookBestPractice := book{
		title:  "La Divina Commedia",
		author: "Dante Alighieri",
		year:   1310,
	}
	_ = bookBestPractice

	// Here, missing fields have zero values
	bookPartial := book{title: "Some new book"}
	fmt.Printf("%#v\n", bookPartial) // main.book{title:"Some new book", author:"", year:0}
}

func readAndUpdateStructs() {
	book1 := book{title: "Anna Karenina"}
	fmt.Printf("%#v\n", book1) // main.book{title:"Anna Karenina", author:"", year:0}

	// Error: undefined
	// pages := book1.pages

	// Access data in struct
	fmt.Println(book1.title) // Anna Karenina

	// Update data in struct
	book1.title = "War and Peace"
	book1.author = "Leo Tolstoy"
	book1.year = 1869
	fmt.Printf("%+v\n", book1) // {title:War and Peace author:Leo Tolstoy year:1869}
	fmt.Printf("%#v\n", book1) // main.book{title:"War and Peace", author:"Leo Tolstoy", year:1869}
}

func comparingStructs() {
	book1 := book{title: "A book", author: "An author", year: 2022}
	book2 := book{title: "A book", author: "An author", year: 2022}
	book3 := book{title: "A book", author: "...An author", year: 2022}

	if book1 == book2 {
		fmt.Println("Book 1 is equal to book 2")
	} else {
		fmt.Println("Book 1 is NOT equal to book 2")
	}
	// Book 1 is equal to book 2

	if book1 == book3 {
		fmt.Println("Book 1 is equal to book 3")
	} else {
		fmt.Println("Book 1 is NOT equal to book 3")
	}
	// Book 1 is NOT equal to book 3
}

func copyingStructs() {
	book1 := book{title: "A book", author: "An author", year: 2022}
	book2 := book1 // <-- This creates a copy!
	if book1 == book2 {
		fmt.Println("Books are equal")
	} else {
		fmt.Println("Books are NOT equal")
	}
	// Books are equal
}

func anonymousStructs() {

	// This is an anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Buffer",
		age:       20,
	}

	// struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Buffer", age:20}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's Age: %d\n", diana.age) // Diana's Age: 20

	// This struct has anonymous fields
	// It can only be assigned without explicit field names
	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 - George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // main.Book{string:"1984 - George Orwell", float64:10.2, bool:false}

	// This is VERY strange
	fmt.Println(b1.string) // 1984 - George Orwell

	// Mixing anonymous and named fields
	type Employee struct {
		name   string
		salary int
		bool
		// bool // Cannot re-declare!
	}

	emp := Employee{"John", 40000, false}
	emp.bool = true
	fmt.Printf("%#v\n", emp) // main.Employee{name:"John", salary:40000, bool:false}
}

func embeddedStructs() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	emp := Employee{
		name:   "John Doe",
		salary: 40000,
		contactInfo: Contact{
			email:   "john.doe@example.com",
			address: "Street 20, London",
			phone:   55512345,
		},
	}

	// {name:John Doe salary:40000 contactInfo:{email:john.doe@example.com address:Street 20, London phone:55512345}}}
	fmt.Printf("%+v\n", emp)
}

func main() {
	// createStructs()
	// readAndUpdateStructs()
	// comparingStructs()
	// copyingStructs()
	// anonymousStructs()
	embeddedStructs()
}
