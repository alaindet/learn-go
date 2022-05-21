package generator

import "os"

var fileContent = `package generated

import "fmt"

type Person struct {
	Name string
	Age int
}

func SomeFunc() {
	fmt.Println("My generated function")
}`

func Generate() {
	_ = os.Mkdir("generated", 0755)
	data := []byte(fileContent)
	_ = os.WriteFile("generated/generated.go", data, 0644)
}
