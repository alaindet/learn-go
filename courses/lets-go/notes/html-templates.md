# HTML Templates

- Go has built-in support for HTML templates via the `html/template` package from the standard library
- You can compose multiple partial templates together
- Templates can have any file extension
- Double curly braces contain Go-specific code that controls the templates, called **actions**
- The most important actions are
  - `define` declares some HTML between `define` and `end` that will later be parsed
  - `template` acts as placeholder for something `define`d later elsewhere
  - `block` like `template`, but declares a default HTML as well
  - `if` Allows conditional rendering of HTML
  - `with` Binds template data to nested fields of the template struct
  - `range` Loops over an iterable field (array, slice, map, channel).

- The actions `if`, `with` and `range` all accepts an optional `else` clause
- `with` and `range` change the value of the **dot** for HTML inside them
  - `with` binds the dot explicitly to what you declare
  - `range` binds the dot to the current element of the loop

- There are also built-in **functions** in the `html/template` package, the most used are
  - `eq` Equals
  - `ne` Not equals
  - `not` Negation
  - `or` Or
  - `index` Extracts a value from an array/slice with an index
  - `printf` Works like `fmt.Sprintf`
  - `len` Length of iterable
  - `:=` Assigns a value to a temporary (template-only) variable
    - Ex.: `{{$foo := len .Something}}`
