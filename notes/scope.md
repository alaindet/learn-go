# Scope

- The scope is the visibility of any given declaration and defines its availability
- i.e. anything declared inside a scope can only be access inside that scope
- Scopes are usually defined via curly braces
- There are only 3 types of scope in Go
  - `File scope`
  - `Package scope`
  - `Block scope`

Ex.:
```go
package main

// File scope: fmt must be imported from another file scope in order to work
// Anything using this file cannot also access the "fmt" package
import "fmt"

// Package scope: this constant is only available in this package
const Done = false

funct main() {
  // Block scope: anything declared here is only available inside this function
  foo := 42
  fmt.Println(foo)
}
```
