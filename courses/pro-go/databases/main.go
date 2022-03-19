package main

import "fmt"

func main() {
	// listDrivers()
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	// rows, err := db.Query("SELECT * FROM Products")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("SELECT * FROM Products", rows)
}
