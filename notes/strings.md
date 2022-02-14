# Strings

Strings are **immutable** slices of bytes in Go. Concatenating strings creates new strings

## Types
Strings are represented by three types in Go
- `string` is the proper string type, is immutable and represented as an array of bytes
- `byte` is an alias for `uint8`, represents an ASCII character, is initialized with a single character in single quotes
- `rune` is an alias for `int32`, represents a Unicode character, is initialized with a single character in single quotes
- Internally, a rune is represented by a *Unicode Code Point*, which is a hexadecimal numeric representation, like `0x61` for "a"
- The default character type is `rune`
- ASCII is a subset of Unicode

## Slicing
- Slicing a string returns a new string from the same backing array
- Slicing returns bytes, not runes
- You can slice by rune (read "letters") and not by byte, by converting the string to an array of runes, then converting again a slice of the new array to string; This is **not efficient** as converting creates a new backing array. Example:

  ```go
  s1 := "答えは質問にあります"
  runeSlice := []rune(s1)             // []int32
  fmt.Println(string(runeSlice[0:3])) // 答えは
  ```

## Raw string
A **raw string** is a string enclosed in backticks which is considered as it is, so no interpolation, trimming, printf verbs, nothing

- New lines are preserved
- No character is escaped
