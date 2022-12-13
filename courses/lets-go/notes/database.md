# Database

- Go exposes a built-in `database/sql` package with a common interface for most SQL database
- `database/sql` needs a database-specific driver to communicate with a real SQL database

## Transactions
- Since `database/sql.Open` returns a pool of connections, transactions are needed to use the same connection for multiple atom operations related together
- Example
  ```go
  db, err := sql.Open(...)
  // Handle error
  defer db.Close()
  tx, err := db.Begin()
  // Handle error
  defer tx.Rollback()
  _, err = tx.Exec("INSERT INTO ...")
  // Handle error
  _, err = tx.Exec("INSERT INTO ...")
  // Handle error
  _, err = tx.Exec("INSERT INTO ...")
  // Handle error
  err = tx.Commit()
  return err
  ```

## Prepared statements
- Allows secure execution of statements by "preparing" a statement on the database, replacing parameters when executing the statement, then discarding the prepared one
- It can have a significant overhead
- For many subsequent executions of the same prepared statement,it's best to call `database/sql.Prepare`
- Example
  ```go
  db, err := sql.Open(...)
  // Handle error
  defer db.Close()
  theStmt, err := db.Prepare("INSERT INTO...") // *sql.Stmt
  // Handle error
  _, err := theStmt.Exec(args...) // Many times
  defer theStmt.Close()
  ```
