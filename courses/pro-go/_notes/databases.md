# Databases

SQL Databases are handled via the basic `database/sql` package, but further specific drivers must be used to communicate with PostgreSQL, MySQL, MariaDB, SQLite etc.

## `sql.Rows.Scan()`

- This method can populate variables with values from a row from a query result (a collection of rows)
- It is very sensitive about order and type of columns in the row, which must match or be compatable with the given variables
- It's not ideal since rows usually map naturally to a struct

Example
```sql
-- MariaDB
CREATE TABLE `Categories` (
  `Id` int(11) NOT NULL,
  `Name` text DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

```go
someDsn := "..." // Etc.
db, err = sql.Open("mysql", someDsn)
_ = err
rows, err := db.Query("SELECT * FROM Categories")

// This is OK
for rows.Next() {
  var id int
  var name string
  rows.Scan(&id, &name)
  // ...
}

// Bad
// What happens is name is not populated (here it's 0) and any subsequent column
// is not populated as well!
for rows.Next() {
  var id string // <-- Id is an int(11), but casting as string is fine here.
  var name int // <-- Name is a text, not compatible with an int!
  rows.Scan(&id, &name)
  // ...
}
```

### Scanning rules from SQL to Go
- SQL strings, numbers and boolean scan to Go equivalent types
- SQL numbers and boolen *can be scanned to strings*
- SQL strings *can be scanned to numbers* if they represent numbers, without overflow
- SQL times can be scanned to Go strings or `*time.Time`
- **Any SQL value** can ba scanned to `*interface{}` and then converted
