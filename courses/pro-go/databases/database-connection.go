package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "progo"
	password = "progo"
	hostname = "127.0.0.1:3306"
	dbname   = "progo"
)

// Thanks to https://golangbot.com/connect-create-db-mysql/
func getDsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Println("Driver:", driver)
	}
}

func connectToDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", getDsn("progo"))
	return
}
