# GO notes

- Any GO program must have a file named `main.go`, with a package and a function inside, both named `main`
  ```go
  // main.go
  package main

  import "fmt"

  func main() {
    fmt.Println("Hello World")
  }

  ```
- To run a program, move into a folder containing a `main.go` file, then run `go run main.go`
