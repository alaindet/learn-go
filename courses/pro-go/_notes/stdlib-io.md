# Input/Output

- The `io` package provides the basics to read from and write to any source, like a stream, a network connection, a file etc.
- The same interfaces are used for any source
- The `io` package provides interfaces, while implementations are distributed in other packages according to the source

## Reader
- Readers are created in various ways across the standard library but they all return an instance of `io.Reader`
- A simple reader for a string can be created with `r := strings.NewReader("foo")`

`io.Reader`
- `Read([]byte)` Read data into `[]byte`, returns count of bytes read and error

## Writers
`io.Writer`
- `Write([]byte)` Writes data from `[]byte`, returns count of bytes written and non-nil error if bytes count is less than the length of `[]byte`
