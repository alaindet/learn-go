# Reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming. It’s also a great source of confusion. [Rob Pyke, The Laws of Reflection](https://go.dev/blog/laws-of-reflection)

- The `reflect.Type` interface defines some methods, one of which is `Kind()`, which returns a `uint` representing a built-in type, e.g. `reflect.Bool` or `reflect.Struct`

- `reflect.PkgPath()` gives the package prefix of some type, empty string for built-in types
  ```go
  package main

  import (
    "reflect"
    "fmt"
  )

  type Product struct {
      Name string
  }

  func GetTypePath(value interface{}) string {
    t := reflect.TypeOf(value)
    path := t.PkgPath()
    if path == "" {
      return "built-in"
    }
    return path
  }

  func main() {
    v1 := "Foo"
    v2 := Product{"Bar"}
    fmt.Printf("(%s) %v\n", GetTypePath(v1), v1) // (built-in) Foo
    fmt.Printf("(%s) %#v\n", GetTypePath(v2), v2) // (main) main.Product{Name:"Bar"}
  }
  ```
