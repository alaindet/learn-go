package main

import (
	"database/sql"
	"fmt"
)

func basicSelectFromDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM Products")

	if err != nil {
		fmt.Println("ERROR: Cannot select from database")
		return
	}

	for rows.Next() {
		var id, category int
		// var id, category string // <-- It's OK too
		var name string
		var price float64
		rows.Scan(&id, &name, &category, &price)
		fmt.Printf(
			// "ID: %d, Name: %s, Category: %d, Price: $%.2f\n",
			"ID: %v, Name: %v, Category: %v, Price: %v\n",
			id,
			name,
			category,
			price,
		)
	}
}

func basicSelectExample() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	basicSelectFromDatabase(db)
}
