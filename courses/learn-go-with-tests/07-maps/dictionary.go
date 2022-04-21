package main

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil // <-- TODO
}

// This works since maps are "kind-of" pointers
// As D. Chaney said "A map value is a pointer to a runtime.hmap structure."
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	if err == nil {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	if err != nil {
		return err
	}

	delete(d, word)
	return nil
}
