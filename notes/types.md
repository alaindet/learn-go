# Types in GO

A type groups values and defines specific methods and operations possible only to those values. Types are divided into
- **Predeclared types** (built-in)
- **Introduced types** (user-defined)
- **Composite types** (user-defined and built-in)

Composite types group together the other types and are divided into
- *array*
- *slice*
- *map*
- *struct*
- *pointer*
- *function*
- *interface*
- *channel*

## Predeclared types
- *int8*, *int16*, *int32*, *int64* (integers)
- *uint8*, *uint16*, *uint32*, *uint64* (positive integers)
- *uint* is an alias for *uint32* or *uint64* based on architecture
- *int* is an alias for *int32* or *int64* based on architecture
- *float32*, *float64* (decimals, can be written without leading 0 or without decimals)
  - Ex.: -.5, 3., .4
- *complex64*, *complex128*
- *byte* is an alias *uint8*
  - Can represent an *ASCII* character
  - Can only be initialized as a single character via single quotes or with a number
  - Ex.: var myByte byte = 'a'
- *rune* is an alias for *int32*
  - It's called a *rune* since *Unicode* symbols are mapped to 32-bit integers
  - Can only be initialized as a single character via single quotes or with a number
  - Ex.: var myRune rune = 'a'
- *bool* true or false
- *string* Unicode chars in double quotes (double quotes only)

## Composite types
- *array* fixed length of elements of the same type
- *slice* variable-length of elements of the same type
- *map* group of elements of the same type indexed by a unique key of any type
  - It is equivalent to Python's dictionary
- *struct* user-defined sequence of *fields* (named elements), each with a name and a type
  - Comparable to classes of other languages
  ```go
  type Car struct {
    brand string
    prince int
  }
  ```
- *pointer* variable storing the memory address of another variable, its zero value is `nil`
  - There is no pointer arithmetic in Go
