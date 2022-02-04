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

## How to run a program
To run a program, move into a folder containing a `main.go` file, then run `go run main.go`

## GO CLI commands
- `go run`: compiles and runs the app, doesn't output an executable
- `go build`: compiles and outputs executable, does not run the app

## Create a module
- `go mod init`
- Then you can run `go build main.go` or `go run main.go` or simply `Ctrl+F5` in VS Code
