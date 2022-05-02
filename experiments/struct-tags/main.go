package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Name  string `firstTag:"The 1st tag of Name" secondTag:"2nd tag"`
		Email string `firstTag:"The 1st tag of Email"`
	}

	u := User{"Bob", "bob@mycompany.com"}
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("\nField: User.%s\n", f.Name)
		fmt.Printf("\tTag value : %q\n", f.Tag)
		fmt.Printf("\tValue of 'firstTag': %q\n", f.Tag.Get("firstTag"))
		fmt.Printf("\tValue of 'secondTag': %q\n", f.Tag.Get("secondTag"))
	}

	/*
		Field: User.Name
			Tag value : "firstTag:\"The 1st tag of Name\" secondTag:\"2nd tag\""
			Value of 'firstTag': "The 1st tag of Name"
			Value of 'secondTag': "2nd tag"

		Field: User.Email
			Tag value : "firstTag:\"The 1st tag of Email\""
			Value of 'firstTag': "The 1st tag of Email"
			Value of 'secondTag': ""
	*/
}
