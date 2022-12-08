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
