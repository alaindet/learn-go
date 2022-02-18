# Structs

- A struct is a sequence of named elements, called **fields**, each having a *name* and a *type*
- It's *similar* (although different) to a *class* as of other programming languages
- Fields are *similar* to instance attributes of class instances
- Structs encapsulate complex data structures, but Go does not have formal support for OOP
- The entire schema of a struct is evaluated at **compile time** and cannot be changed at run time

## Comparison

- Structs can be compared
- Structs are equal if all fields are individually equal to each other

## Anonymous structs and fields
- An anonymous struct can be assigned to a variable and be used
- An anonymous field is named after its type, meaning that **a struct cannot have more than one anonymous field of the same type**
  ```go
  type hello struct {
    name string
    salary int
    bool // This will have "bool" name
    // bool // Cannot re-declare bool!
    int // This will have "int" name
  }
  ```
