package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "events.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)
	DB = db
	createTables()
}

func createTables() {

	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS "events" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" TEXT NOT NULL,
			"description" TEXT NOT NULL,
			"location" TEXT NOT NULL,
			"date_time" DATETIME NOT NULL,
			"user_id" INTEGER
		)
	`)

	if err != nil {
		panic("Could not create \"events\" table")
	}
}
