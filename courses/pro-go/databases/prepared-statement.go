package main

import (
	"database/sql"
	"fmt"
)

func createCategory(db *sql.DB, name string) (int64, error) {
	createStmt, err := db.Prepare("INSERT INTO Categories (Name) VALUES (?)")

	if err != nil {
		error := fmt.Errorf("[ERROR] Cannot prepare statement: %s", err.Error())
		return 0, error
	}

	result, err := createStmt.Exec(name)

	if err != nil {
		error := fmt.Errorf("[ERROR] Cannot create category: %s", err.Error())
		return 0, error
	}

	id, err := result.LastInsertId()

	if err != nil {
		error := fmt.Errorf("[ERROR] Cannot get last inserted ID: %s", err.Error())
		return 0, error
	}

	return id, nil
}

func updateProductCategory(
	db *sql.DB,
	categoryID int64,
	productIDs ...int,
) (int64, error) {
	categoryIDAsInt := int(categoryID)
	var rowsAffected int64

	// WARNING: This is inefficient as the statement is executed FOR EACH product!
	// TODO: This can be done in a single statement with IN(...)
	updateStmt, err := db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")

	if err != nil {
		error := fmt.Errorf("[ERROR] Cannot prepare statement: %s", err.Error())
		return 0, error
	}

	for _, productID := range productIDs {
		result, err := updateStmt.Exec(categoryIDAsInt, productID)

		if err != nil {
			continue
		}

		count, err := result.RowsAffected()
		_ = err

		if err != nil {
			continue
		}

		rowsAffected += count
	}

	return rowsAffected, nil
}

func preparedStatementExample() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	categoryId, err := createCategory(db, "Miscellaneous")

	if err != nil {
		panic(err)
	}

	fmt.Println("Created category with ID:", categoryId)

	productsUpdatedCount, err := updateProductCategory(db, categoryId, 1, 2)

	if err != nil {
		panic(err)
	}

	fmt.Println("Rows affected:", productsUpdatedCount)
	fmt.Printf("Products #1 and #2 have been added to category #%d\n", categoryId)
}
