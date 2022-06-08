package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	dbMaxTries = 10
	dbRetryIn  = 1000 // milliseconds
	webPort    = "80"
)

var db *sql.DB

func main() {
	db := initDB()
	fmt.Println("db", db)
}

// TODO: Singleton with sync.Once()?
func initDB() *sql.DB {
	if db == nil {
		conn, err := connectToDB()

		if err != nil {
			log.Panic("cannot connect to database")
		}

		db = conn
	}

	return db
}

func connectToDB() (*sql.DB, error) {
	tries := 0
	dns := os.Getenv("DSN")

	for {
		conn, err := openDB(dns)

		if tries > dbMaxTries {
			return nil, fmt.Errorf("cannot connect to the database")
		}

		if err != nil {
			log.Println("database not ready yet")
			tries++
			time.Sleep(time.Duration(dbRetryIn) * time.Millisecond)
			continue
		}

		log.Println("connected to database")
		return conn, nil
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
