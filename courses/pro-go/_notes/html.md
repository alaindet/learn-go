# HTML

- Go has several built-in template engines, one of which is `html/template` and is used to produce HTML based on templates
- HTML templates are simple .html files containing **actions**
- Templates must be loaded and parsed as a first step, then they are *executed*, meaning a parsed template receives actual data to be combined with the template to generate some HTML output

## Actions

Actions in HTML template files are included in **double curly braces** and allow for any kind of interpolation: simple values, expressions, piping, function calls, conditionals, nesting, looping etc.

Examples
- `{{ val }}`
- `{{ expr }}`
- `{{ val.SomeField }}`
- `{{ val.Method arg }}`
- `{{ func arg }}`
- `{{ expr | val.Method }}`
- `{{ expr | func }}`
- `{{ range val }} <other HTML> {{ end }}`
- `{{ range val }} <other HTML> {{ else }} <other HTML> {{ end }}`
- `{{ if expr }} <other HTML> {{ else if expr }} <other HTML> {{ else }} <other HTML> {{ end }}`
- `{{ with expr }} <other HTML> {{ end }}`
- `{{ define "aTemplateName" }} <other HTML> {{ end }}`
- `{{ template "aTemplateName" expr }}`
- `{{ block "aTemplateName" expr }} <other HTML> {{ end }}`

### Built-in
Apart from user-defined functions, there are some built-in functions in Go to format values and evaluate expressions in templates

Formatting
- `print`
- `printf`
- `println`
- `html`
- `js`
- `urlquery`
- `slice`
- `index`
- `len`

Evaluating expressions
- `eq a b` (equal) True if a == b
- `ne a b` (not equal) True if a != b
- `lt a b` (less than) True if a < b
- `le a b` (less or equal to) True if a <= b
- `gt a b` (greater than) True if a > b
- `ge a b` (greater or equal to) True if a >= b
- `and a b` True if a == true and b == true
- `not a` True if a != true
