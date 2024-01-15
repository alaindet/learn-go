package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// gobFileExample()
	gobRuntimeExample()
}

func gobRuntimeExample() {
	people := []Person{
		{Name: "Alice", Age: 80},
		{Name: "Bob", Age: 90},
		{Name: "Charlie", Age: 100},
	}

	serializedBin, err := ToGob[[]Person](people)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	deserializedPeople, err := FromGob[[]Person](serializedBin)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Printf("Deserialized data: %+v\n", deserializedPeople[1])
}

func gobFileExample() {
	person := Person{
		Name: "Alice",
		Age:  80,
	}

	cacheFile := "data.gob"

	if err := ToGobFile(person, cacheFile); err != nil {
		fmt.Println("ERROR", err)
		return
	}

	deserializedPerson, err := FromGobFile[Person](cacheFile)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Printf("Deserialized data: %+v\n", deserializedPerson)
}

func ToGob[T any](data T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func ToGobFile[T any](data T, path string) error {
	databin, err := ToGob(data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, databin, 0644); err != nil {
		return err
	}

	return nil
}

func FromGob[T any](databin []byte) (T, error) {
	buffer := bytes.NewBuffer(databin)
	decoder := gob.NewDecoder(buffer)
	var data T

	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func FromGobFile[T any](path string) (T, error) {
	var data T

	databin, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}

	return FromGob[T](databin)
}
