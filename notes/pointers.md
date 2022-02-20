# Pointers

## Introduction
- RAM is conceptually a sequence of adjacent boxes or cells in line
- A cell has a *unique address* expressed as an hexadecimal number, like 0xC000A16120
- The address is called **memory location**
- Each cell is a single value, CPU "simply" gets and sets values into cells
- GO does not allow any pointer arithmetic, unlike C
- A function can return a pointer, unlike C

## Pointers and variables
- A *variable* is a label for a memory location
- For example, the instruction `var a int = 42` puts the `int` value `42` into a memory location and assigns it a label of `a`
- A **pointer** is a variables storing the memory address of another variable!
- A non-initialized pointer has value `nil`

## Initialization and values
- Pointers can **mutate** the value they are referring to
- Pointers can be initialized using the **address operator** ampersand (`&`) if they reference an existing variable
  ```go
  name := "John"
	addr := &name
	fmt.Println(name, addr) // John 0xc000010230
  ```
- The type of a pointer depends on the type of the references value and is preceded by `*`
- A type like `*int` is read as "pointer of int"
  ```go
  var x int = 2
	ptr := &x
	fmt.Printf("%T %v", ptr, ptr) // *int 0xc000014080
  ```
- Pointers can be initialized without any value (default to `nil`) like this
  ```go
  var myPointer *int // It is <nil>
  ```

## Dereferencing operator
- If you want to use the value to which a pointer refers, use the **dereferencing operator** `*`
  ```go
  x := 42
  p := &x
  y := *p // This assigns the **VALUE** of x to y, copying it
  fmt.Println(y) // 42
  y = 45 // This changes the value of y, which, is not related to x anymore
  *p = 47 // This changes the value of the referenced variable of p, which is x
  fmt.Println(x, y) // 47 45
  ```

- In short
  - `&value => pointer` if you have a value you turn it into an address or pointer by using the ampersand operator
  - `*pointer => value` and if you have pointer you turn it into value value by using the star operator

## Star symbol `*`
- The star symbol `*` can be confusing when dealing with pointers
  ```go
  x := 1
  // Here, it means p will reference a variable of type int
  var p *int = &x
  // Here, you are dereferencing p to use x's value directly
  y := *p
  ```

## Comparison
- Pointers can reference to other pointers and be compared to anything
- Pointers are equal if they point to the same variable (they have the same address) or if they are both `nil`

## Pointers and functions
- There is no concept of *pass by reference* in Go since every argument is copied
- However, passing a pointer as argument allows the function to modify the "referenced" value
- This is improperly defined as **pass by reference** in Go
- Please note that a pointer argument is still a copy of the given pointer

### Pass by value or by reference?
- In general, avoid passing any array to functions as they're copied
- Passing by value (default) is cheaper and easier, but can end up taking much memory
- Passing by reference puts pressure on the garbage collector, but it does not occupy other memory for copying arguments
- Start with the default, and pass by reference only if necessary
