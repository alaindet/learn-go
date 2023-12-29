# Enums
- Explicit enumeration does not exist in Go
- Several patterns can be used to simulate enumeration, like
  - A custom type, constants defined with `iota` and a zero value for the enum
  - A custom type, structs and runtime variables
  - TODO: Code generation?

## Constants
- A new type `MyEnum` (replace with your name) is defined by aliasing `int` built-in type
- Constants are declared of type `MyEnum` via the `iota` keyword
- Constants have names prefixed with `MyEnum` to guarantee scoping, ex.: `MyEnumRed`
- An optional `MyEnumUnknown` constant could be defined as the first constant, so that `iota == 0`; this constant can be used as the *zero value* of `MyEnum`

## Structs
- A new type `MyEnum` (replace with your name) is defined with a struct with only a `name` field
- Enums values are declared as runtime variables each containing a `MyEnum` struct with a unique string for the `name` field

## Resources
- https://threedots.tech/post/safer-enums-in-go/
- https://www.mohitkhare.com/blog/enums-golang/
