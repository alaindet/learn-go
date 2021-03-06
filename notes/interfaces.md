# Interfaces
- An interface is a collection of method signatures that an object (mostly named types) can implement
- Interfaces define how the object behave
- They are **NOT** equivalent to generics
- Example declaration is
  ```go
  type shape interface {
    area() float64
    perimeter() float64
  }
  ```
- Implementation of interfaces is **implicit**, so if a named type has all the methods described by the interface, it implicitly implements it
- There's no explicit keyword `implements` as in other languages
- Conceptually, this is **duck typing**
- The zero value of interfaces is `nil`
- Interfaces values are a pair of a **concrete value** and a **dynamic type**

- A variable containing an interface has *two types* at the same time, a **static type** and a **dynamic type**
- The *static type* is interface type per se and cannot be changed
- The *dynamic type* is the type implementing the interface
- An interface "absorbs" the dynamic value and type of the value you assign it
  ```go
  // ... Previously defined "shape" interface, "circle" type, "rectangle" type
  var s shape
  fmt.Printf("%T\n", s) // <nil>
  ball := circle{radius: 10.0} // This is a concrete value
  s = ball
  // Type of s chaged to the "dynamic type" main.circle
  fmt.Printf("%T\n", s) // main.circle
  frame := rectangle{width: 16.0, height: 9.0} // This is a concrete value
  s = frame
  // Type of s changed again!
  fmt.Printf("%T\n", s) // main.rectangle
  ```

## Polymorphism
- Interfaces allow for polymorphism in Go
- *Polymorphism* is the concept for which two objects of different types can perform the same operation regardless of the implementation details, if they share an interface
- Basically, an interface is a contract objects adhere to so that if an object implements an interface, it *must* have what the interface prescribes

## Type Assertion
- It can only be performed on interfaces
- It provides access to an interface's dynamic type
```go
// Imported "log"
// Declared shape interface, circle type, volume method on circle
var s shape = circle{radius: 24.5}
c, ok := s.(circle) // <-- This is a type assertion!
if !ok {
		log.Fatal("Error")
}
volume := c.volume()
```

## Type Switch
- A type switch is very similar to a regular `switch` case, but it works on type assertions using the special `.(type)` notation that returns the type name of a variable
```go
var s shape = circle{radius: 24.5}

switch value := s.(type) {
case circle:
    fmt.Printf("%#v is a circle\n", value)
case rectangle:
    fmt.Printf("%#v is a rectangle\n", value)
}
// main.circle{radius:24.5} is a circle
```

## Embedding
- GO explicitly does not support inheritance
- Interfaces can be **embedded** to extend another interface, it happens at compile time
- Interfaces creating a circular embedding circle will not compile

## The empty interface
- This is single-handedly the most dangerous feature of GO
- Please avoid using it or use with extreme caution
- Since interface implementation is implicit, **ANY** named type implements the empty interface `interface{}`, so coding against `iunterface{}` means the type can be anything!
- The empty interface is equivalent to TypeScript's `any` type
- The empty interface **moves typing at run time**
- Example, this is all valid!!!
  ```go
  type person struct {
    info interface{}
  }

  p := person{info:"Hello World"}
  p.info = 42
  p.info = []string{"I", "am", "Iron", "Man"}
  p.info = p
  p.info = [4]int{1, 2, 3, 4}
  ```

## Implementation
- Methods implementing an interface must only match these
  - Same name
  - Same parameters types
  - Same result types
- This means methods can have **different parameter names and named results**
