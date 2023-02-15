package main

import (
	"errors"
	"fmt"
)

func main() {
	err := Try(func(try func(err error)) {
		fmt.Println("This is a dangerous function")
		fmt.Println("This thing could break any minute")
		try(ok("Will it break now?"))
		try(ok("What about now?"))
		try(ko("BOOM!")) // <-- Look at this!
		try(ok("This will never run"))
		try(ok("This won't run either"))
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Keep going...")
}

func ok(message string) error {
	fmt.Println(message)
	return nil
}

func ko(message string) error {
	return errors.New(message)
}
