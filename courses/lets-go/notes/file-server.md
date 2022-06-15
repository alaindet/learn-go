# File server

- Go has a built-in file server `http.FileServer` in `net/http` standard package
- It plays nice with an existing servermux
- It sanitizes relative paths to mitigate attacks
- It supports range requests for large files
- Reduces latency by sending 304 Not Modified status to clients requesting the same file multiple times
- The `Content-Type` is automatically set based on file extension or custom rules
- Serving a file from a custom HTTP handler is done via `http.ServeFile(w http.ResponseWriter, r *http.Request, path string)`
