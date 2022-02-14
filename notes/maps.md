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

## Comparison
- Maps can only be compared to `nil`
- Two maps can be compared using their string representation with `fmt.Sprintf()`
  ```go
  a := map[string]string{"foo": "bar"}
	b := map[string]string{"foo": "bar"}
  s1 := fmt.Sprintf("%s", a) // map[foo:bar]
  s2 := fmt.Sprintf("%s", b) // map[foo:bar]
  fmt.Println(s1 == s2) // true
  ```
- `fmt.Sprintf()` sorts the key-value pairs by key before build the string, so that two string represenations of two maps with the same key-value pairs expect order are equal
  ```go
  c := map[string]string{"a": "b", "c": "d"}
	d := map[string]string{"c": "d", "a": "b"}
	s3 := fmt.Sprintf("%s", c)
	s4 := fmt.Sprintf("%s", d)
	fmt.Println(s3, s4)                      // map[a:b c:d] map[a:b c:d]
	fmt.Println("Are maps equal?", s3 == s4) // Are maps equal? true
  ```

## Cloning

- A map is internally represented by a data structure called **map header**, which holds the key-value pair data
- Assigning a map to another map just creates *a new reference* to the same underlying *map header*
