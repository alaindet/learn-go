# HTML Templates

- Go has built-in support for HTML templates via the `html/template` package from the standard library
- You can compose multiple partial templates together
- Templates can have any file extension
- Double curly braces contain Go-specific code that controls the templates, called **actions**
- The most important actions are `define`, `template` and `block`

## `define`
- `define` declares a portion of HTML that will later be parsed and used from Go code
- The HTML portion must live between `define` and `end`

## `template`
- `template` is like `require` in PHP and imports something previously defined with `define`
- Nothing is printed if the template is not defined


## `block`
- `block` imports a HTML portion like `template`, but accepts a HTML portion to be used as default when importing a template which is not yet `define`d
- The HTML portion must live between `block` and `end`
- You can use it as `template` without a default HTML portion so that it acts like an optional template
  - Example
  ```
  {{block "sidebar" .}}{{end}}
  ```
- It's usually preferred instead of `template` so that you can later provide a default HTML portion
