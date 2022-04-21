package main

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not found")

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

// This works since maps are "kind-of" pointers
// As D. Chaney said "A map value is a pointer to a runtime.hmap structure."
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
