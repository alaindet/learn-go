# Maps

- It's a composite type storing key-value pairs
- Adding, accessing and deleting key-value pairs takes constant time
- Keys and values must have the same type declared ahead of time
- Keys can have any **comparable** type (anything comparable with `==`), not just strings and numbers (like JavaScript)
- Keys must be unique
- Floats are discouraged as keys since comparison has known issues of precision
- Maps cannot be compared to each other, but can only be compared to `nil`
- The zero value of a map is `nil`
- Maps are considered **unordered** data structures in Go
