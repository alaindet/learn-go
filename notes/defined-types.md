# Defined types

- A *defined type* is a *named type* defined by the user from an existing type called *source type*, be it another *defined type* or a built-in type
- A *defined type* has a specific name and can have new methods
- The *source type* gives representation, operations and byte size to the defined type
- Defined and source types share representations etc., but *they are completely different*
- Operations between the source and the defined type need type conversion
- Type conversions are not always possible between types
- There is **not type-hierarchy in Go**
