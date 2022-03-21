package main

import (
	"database/sql"
	"fmt"
)

func selectSingleProduct(db *sql.DB, productId int) (p Product, error error) {
	row := db.QueryRow(`
		SELECT
			p.Id as product_id,
			p.Name as product_name,
			p.Price as product_price,
			c.Id as category_id,
			c.Name as category_name
		FROM
			Products p
			INNER JOIN Categories c ON p.Category = c.Id
		WHERE
			p.Id = ?
	`, productId)

	if row.Err() != nil {
		error = fmt.Errorf("[ERROR] Product not found: %s", row.Err().Error())
		return
	}

	err := row.Scan(
		&p.Id,
		&p.Name,
		&p.Price,
		&p.Category.Id,
		&p.Category.Name,
	)

	if err != nil && err != sql.ErrNoRows {
		error = fmt.Errorf("[ERROR] Database error: %s", err.Error())
		return
	}

	if err == sql.ErrNoRows {
		error = fmt.Errorf("[ERROR] Product not found: %s", err.Error())
		return
	}

	return
}

func selectSingleRowsExample() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	for _, id := range []int{1, 3, 10} {
		product, err := selectSingleProduct(db, id)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(product)
	}
}
