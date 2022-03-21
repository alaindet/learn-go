package main

import (
	"database/sql"
	"fmt"
)

func insertProduct(db *sql.DB, p *Product) (int64, error) {
	result, err := db.Exec(
		`INSERT INTO Products (Name, Category, Price) VALUES (?, ?, ?)`,
		p.Name,
		p.Category.Id,
		p.Price,
	)

	if err != nil {
		error := fmt.Errorf("[ERROR] Cannot create product: %s", err.Error())
		return 0, error
	}

	id, err := result.LastInsertId()

	if err != nil {
		error := fmt.Errorf("[ERROR] No new ID was created for product: %s", err.Error())
		return 0, error
	}

	return id, nil
}

func insertProductExample() {
	db, err := connectToDatabase()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()
	fmt.Println("Connected to database")

	newProduct := Product{
		Name: "Stadium",
		Category: Category{
			Id: 2,
		},
		Price: 79500,
	}

	productId, err := insertProduct(db, &newProduct)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// selectSingleProduct => select-single-rows.go
	product, err := selectSingleProduct(db, int(productId))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(product)
}
