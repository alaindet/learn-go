/*
A struct can be serialized and unserialized to create a copy

In Go, binary serialization outputs a "gob" file, which represents any built-in
and custom Go type as binary
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (p *Person) DeepCopyWithSerialization() *Person {
	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)
	_ = enc.Encode(p)

	dec := gob.NewDecoder(buffer)
	newPerson := &Person{}
	_ = dec.Decode(newPerson)

	return newPerson
}

func serializationExample() {
	john := NewPerson(
		"John",
		&Address{"123 Foo Street", "London", "UK"},
		[]string{"Chris", "Matt"},
	)

	jane := john.DeepCopyWithSerialization()
	jane.Name = "Jane"

	fmt.Println("serializationExample")
	printPerson(jane)
}
