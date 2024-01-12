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
	person := Person{
		Name: "Alice",
		Age:  80,
	}

	cacheFile := "data.gob"

	if err := ToGobFile(person, cacheFile); err != nil {
		fmt.Println("ERROR", err)
	}

	deserializedPerson, err := FromGobFile[Person](cacheFile)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	fmt.Printf("Deserialized data: %+v\n", deserializedPerson)
}

func ToGob(data any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func ToGobFile(data any, path string) error {
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
