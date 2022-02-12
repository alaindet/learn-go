# Arrays and Slices

| Arrays                              | Slices                                        |
| ----------------------------------- | --------------------------------------------- |
| Fixed length                        | Dynamic length                                |
| Length is part of its type          | Length is not part of its type and            |
| Length is defined at compile time   | Length is defined at run time, can be changed |
| Uninizialited array has zero-values | Is equal to `nil` when uninitialized          |
| Elements must be of same type       | Elements must be of same type                 |
| Can create a keyed array            | Can create a keyed slice                      |

## `nil`
Nil (from latin *nihil/nil*) value means the variable is not defined yet, it's treated more like a default value rather than an explicit missing value. In JavaScript, `undefined` is the equivalent, while `null` is the explicit absence of a value. There is not Go equivalent for `null`

## Comparison

### Arrays
Arrays can be compared and are equal if
- they have the same length
- they have the same elements
- elements are in the same order

### Slices
Slices cannot be compared, a custom for loop or a function is needed

## Backing Array
- Every slice has a **backing array** or **underlying array** completely managed by Go
- It stores elements
- The internal data structure of a slice is called **slice header**
- The slice header has 3 fields
  - the **address** of the backing array (pointer)
  - The **length** of the slice, returned in code by `len()`
  - The **capacity** of the slice, which is the length of the underlying array, returned in code by `cap()`
- The slice header is a runtime representation of a slice
- The **nil slice** is a slice without a backing array

- Any *slice expression* (e.g. `s[0:4]`) returns a new slice based on the same *backing array*
