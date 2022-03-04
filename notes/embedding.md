# Embedding

- Embedding is the mechanism with which Go achieves **composition** and substitutes *inheritance* of other common languages
- Embedding is performed on a struct or an interface, which gets embedded inside an **enclosing** type at compile time
- Since embedding is not inheritance, an enclosing type cannot be used instead of its embedded type, where the embedded type is required. Ex.: if `Bar` encloses `Foo` and you declare a function accepting an argument of type `Foo`, you cannot use structs of type `Bar` instead

## Structs
- The fields of the embedded type can be accessed though they were declared in the *enclosing type*, through a mechanism called **field promotion**
- *field promotion* does not work when initializing with a literal syntax
- *field promotion* is not used if the enclosing type has already fields with the same name as the embedded type (the enclosing one's takes precedence)
- If a struct embeds two or more types sharing one or more fields with the same name, trying to access those fields generates an *ambiguous selector* fatal error

Example
```go
type Foo struct {
    Name  string
    Price float64
}

func (f *Foo) Print() string {
	return fmt.Sprintf("Name: %q, Price: %.2f", f.Name, f.Price)
}

type Bar struct {
    *Foo
    Name     string
    Category string
}

func main() {

	// This works, but it is not recommended
    // b := &Bar{
	// 	&Foo{"My Foo", 123.00},
	// 	"My Bar",
	// 	"My Bar Category",
	// }

	// This is better, best is using constructors
	f := &Foo{"My Foo", 123.00}
	b := &Bar{f, "My Bar", "My Bar Category"}

    fmt.Println(b.Name)     // My Bar
    fmt.Println(b.Foo.Name) // My Foo
    fmt.Println(b.Price)    // 123.00
	b.Name = "Changed name"
	b.Price = 321.00
	fmt.Println(b.Print()) // Name: "My Foo", Price: 321.00
}
```

Please note:
- `b.Name` access the `Name` field of `Bar` since it exists
- `b.Foo.Name` explicitly accesses the `Name` field of the embedded type `Foo`
- `b.Price` uses the **field promotion** and accesses the `Price` field of `Foo` since no `Price` field is explicitly declared in the enclosing type `Bar`

### Methods
- Methods defining an embedded struct as a receiver are accessible from the enclosing type too

Example
```go
type Foo struct {
    Name  string
    Price float64
}

func (f *Foo) Print() string {
	return fmt.Sprintf("Name: %q, Price: %.2f", f.Name, f.Price)
}

type Bar struct {
    *Foo
    Name     string
    Category string
}

func main() {
    f := &Foo{"My Foo", 123.00}
	b := &Bar{f, "My Bar", "My Bar Category"}
	// Note: Despite Bar having the Name field, this returns "My Foo" since
	// the method expects a Foo anyway
	fmt.Println(b.Print()) // Name: "My Foo", Price: 123.00
}
```
