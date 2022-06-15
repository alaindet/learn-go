# Servermux

- A servermux for Go is just a regular HTTP router
- All incoming HTTP requests are served in their own goroutine
- **WARNING**: Servermux cannot route based on HTTP method
- **WARNING**: Servermux cannot match URL variables
- **WARNING**: Servermux cannot match URL patterns expressed as regexp
- URL patterns provided to the servermux only match when the request URL **exactly matches** a fixed path
- Order of registration of routes does not matter since **longer paths take precedence**
- URLs are always sanitized before matching, ex.: `/foo/bar/..//baz` => `/foo/baz`
- If you provide host-specific patterns, those are matched first if possible
  - Ex.:
  ```go
  mux.HandleFunc("/", snippetView) // Second
  mux.HandleFunc("example.org/", snippetView) // First
  ```
- **URL patterns** are divided into two categories, fixed paths and subtree paths

## Fixed paths

Ex.: `/snippet/view`

Or simply "paths", those are paths with **no trailing slash**, they match only when the request URL matches **exactly** with them

## Subtree paths

Ex.: `/` or `/assets/`

Or simply "path prefixes", those are paths with **a trailing slash** and they match anything that is **prefixed by** them, ex.: `/assets/` matches `/assets`, `/assets/logo.jpg` and `/assets/icons/64x64/admin.svg`, while `/` catches all routes

## HTTP Headers

- Go automatically populates these headers in the response
  - `Date`
  - `Content-Length`
  - `Content-Type` This is performed via `http.DetectContentType()` and if no valid content type is detected, `Content-Type: application/octet-stream` is used as default

- **WARNING:** `http.DetectContentType()` does not distinguish JSON from plain text, so manual header setting is requested
  ```go
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"name":"Alain"}`))
  ```
