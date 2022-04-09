package main

import (
	"fmt"

	"gophercises-cyoa/cyoaweb"
)

func main() {
	s := cyoaweb.LoadStory()
	h := cyoaweb.NewHttpHandler(s)
	fmt.Println(h)
}
