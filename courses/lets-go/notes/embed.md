# Embedding

- `embed` is a standard library package available since Go 1.16
- It "embeds" any file by putting it into the final binary you build with `go build`
- It allows to embed and minimize dependencies so you can have a single binary file
- Performance-wise, embedding small frequently used files grants the benefit of not reading them from disk everytime, but large or too many files should not be embedded
- Good use cases: Loading settings, templates, fonts

## How
- Embedding is flagged via code through special comments called **comment directives**, which match the pattern `//go:embed <paths>`, for example:
  - Ex.: This imports all files from the "templates" folder into the `templates` variable
    ```go
    import "embed"
  
    //go:embed "templates"
    var templates embed.FS
    ```
- Embedded files can only be cast to these three types
  - `string`
  - `[]byte`
  - `embed.FS`
  - Ex.:
    ```go
    import "embed"

    // Single file content read as text
    //go:embed manual.txt
    var man string

    // Single file content read as slice of bytes
    //go:embed logo.png
    var logo []byte

    // Directory, allows to access multiple files
    //go:embed templates/*
    var templates embed.FS
    ```
- You can embed multiple paths if you use the `embed.FS` type
  - Ex.:
  ```go
  import "embed"

  //go:embed "html" "static"
  var theFiles embed.FS
  ```
- Paths are calculated relative to the file with the comment directive
- Paths **cannot have** `.`, `..`, leading `/` and trailing `/`, meaning you can only embed files from the file's folder or subfolders
- Paths that are directories embed **ALL** files in it, expect files starting with `.` and `_` (conventionally hidden files)
- Paths prepended with `all` load all files instead. Ex.: `//go:embed all:mydir`
- Windows-specific `\` dir separator is not used, use `/` on any system
- The root of the `embed.FS` filesystem is the folder containing the `.go` file with the comment directive
