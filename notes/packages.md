# Packages
- A package is a folder containing many `.go` files having the same **package statement** at the beginning
- All `.go` files in a folder MUST belong to the same package
- Types of packages are
  - **executable packages**
    - generates executables files
    - predefined package name must be `main`
  - **non-executable packages** (libraries, dependencies)
    - used by other libraries or applications
    - can have any package name
    - can't be executed, only imported


## Access Control
- Anything declared in a package, in any file of the package, is considered to be **private** by default
- Something "private" can still be accessed by anything else from inside the package, but not from outside
- Anything **starting with a capital letter** in a package is **implicitly considered to be public**, which means it can be used from any external code *importing* the package
- Access control rules (package prefix and private fields) do not apply for code inside a package


## Importing
- A user-define package must be imported with a string including at least the module name and the package name (ex.: `"packages/store"`)
- A package can have nested packages which are simply nested folders with `.go` files declaring the nested package as their namespace
- Imported declarations (structs, methods, interfaces, functions etc.) must be prefixed with the package name (ex.: `store.Product` struct)
- An import can define an **alias** in order to avoid conflicts with other packages
- A package defines a namespace, so that there cannot be two variables, types, structs, functions or interfaces with the same name in a package
- The **dot import** is performed by using a dot `.` as an alias and exposes any declaration from the aliased package

Examples
```go
import "packages/store" // Import used-defined package
import "packages/store/cart" // Import a nested package
import myFmt "packages/fmt" // Give an alias to the package since "fmt" already exists
import . "packages/fmt" // Import everything from the package and use it without prefix
```


## Initialization
- Any `.go` file can perform some **initialization** function to perform some tasks/computations once
- Initialization is performed after all imports, constants and variables have been evaluated
- Initialization is done via the `init()` function, which is the only function name with special meaning other than `main()` in Go
- `init()` has no arguments and no result values
- `init()` cannot be invoked programmatically
- A single file can declare two or more `init()` functions, although NOT RECOMMENDED
- Any file in a package can have its own `init()` function
- Multiple `init()` functions in multiple files in a package are executed following the **alphabetical order** of the file names, but it is NOT RECOMMENDED to write `init()` functions relying on each other since order is not guaranteed from the Go specification
