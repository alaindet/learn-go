package main

import (
	"database/sql"
	"fmt"
)

func selectProductsFromDatabaseWithPlaceholders(db *sql.DB) []Product {
	products := []Product{}
	sqlStatementWithPlaceholders := `
		SELECT Id, Name, Price
		FROM Products
		WHERE Price < ?
	`
	placeholder1Value := 100
	rows, err := db.Query(sqlStatementWithPlaceholders, placeholder1Value)

	if err != nil {
		panic("ERROR: Cannot select from database")
	}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Price)

		if err != nil {
			fmt.Println("Scan error", err.Error())
			break
		}

		products = append(products, p)
	}

	return products
}

func preparedStatementExample() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	products := selectProductsFromDatabaseWithPlaceholders(db)
	fmt.Println(products)
}
