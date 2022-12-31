# Logs

- Logging allows to trace relevant parts of the execution of code
- Logs in every language and library tend to have **levels** (3-7 levels usually)
- Standard levels include
  - **TRACE** Tracks the normal execution of code, sometimes optional
  - **DEBUG** Used to provide diagnostics on what's happening (and why's not working)
  - **INFO** Neutral notification about what usually happens
  - **WARNING** Signals something worth the attetion, but it does not stop the execution
  - **ERROR** Signals something critical, unexpected that should be addressed, does not stop execution
  - **FATAL** or **PANIC** Worst possible case, execution MUST stop

- Go has a built-in `log` package that allows to create custom loggers
- Loggers created with `log.New()` are concurrency-safe
- Loggers that can write to stdout, files and remote services
- Loggers can also provide custom templates and additional info you can pick from (date, time, file name and line that called it)

- **Convention**: It's best to call `Panic()` and `Fatal()` only from `main()` and return normal errors from inner code
- **Convention**: It's best to output to STDOUT and STDERR by default and then optionally write to some other stream/file with something like this
  ```
  $ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
  ```
