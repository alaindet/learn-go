# Data

## Site data

- Site data is available in layouts and themes via the `.Site` key, ex.: `{{ .Site.Title }}`
- Site data is data declared in `/config.toml`
- Example
  ```
  // config.toml

  [params]
    author = "John Smith"
    description = "My portfolio website demo"
  ```

  ```
  // layouts/_default/index.html

  <meta name="author" content="{{ .Site.Params.author }}">
  <meta name="description" content="{{ .Site.Params.author }}">
  ```

## `/data/`
- The `/data/` folder contains `JSON`, `YAML` and `TOML` files that can only be accessed by layouts
