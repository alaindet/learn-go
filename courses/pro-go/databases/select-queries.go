package main

import (
	"database/sql"
	"fmt"
)

func selectProductsFromDatabase(db *sql.DB) []Product {
	products := []Product{}
	rows, err := db.Query(`
		SELECT
			p.Id as product_id,
			c.Id as category_id,
			c.Name as category_name,
			p.Name as product_name,
			p.Price as product_price
		FROM
			Products p
			INNER JOIN Categories c ON p.Category = c.Id
	`)

	if err != nil {
		panic("ERROR: Cannot select from database")
	}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(
			&p.Id,
			&p.Category.Id,
			&p.Category.Name,
			&p.Name,
			&p.Price,
		)

		if err != nil {
			fmt.Println("Scan error", err.Error())
			break
		}

		products = append(products, p)
	}

	return products
}

func selectQueriesExamples() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	products := selectProductsFromDatabase(db)
	fmt.Println(products)
}
