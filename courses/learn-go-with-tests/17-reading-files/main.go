package main

import (
	"log"
	"os"

	"learn_go_with_tests/reading_files/blogpost"
)

func main() {
	posts, err := blogpost.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		log.Printf("%+v\n", p)
	}
}
