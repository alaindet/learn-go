# Defined types

- A *defined type* is a *named type* defined by the user from an existing type called *source type*, be it another *defined type* or a built-in type
- A *defined type* has a specific name and can have new methods
- The *source type* gives representation, operations and byte size to the defined type
- Defined and source types share representations etc., but *they are completely different*
- Operations between the source and the defined type need type conversion
- Type conversions are not always possible between types
- There is **not type-hierarchy in Go**

## Alias types

An *alias type declaration* looks like this

```go
type T1 = T2
// Normal declaration is
// type T1 T2
```

- The aliased type just **gives a new name to the same referred type**
- Ex.: The expression `type T1 = T2` assignes the new name `T1` to the existing type `T2` so that `T1` and `T2` point to the same type and *are considered the same type*!
- There is no need for type conversions in an operation with two aliased types
- Avoid using aliased types whenever possible

- There are built-in examples of aliased types, like
  - `byte` and `uint8`
  - `rune` and `int32`
