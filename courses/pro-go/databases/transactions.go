package main

import (
	"database/sql"
	"fmt"
)

/*
	Create a new category and associate it with given products
*/
func createAndUseCategory(
	db *sql.DB,
	categoryName string,
	productIDs ...int,
) (err error) {
	tx, err := db.Begin()

	if err != nil {
		return fmt.Errorf("[ERROR] Cannot begin transaction")
	}

	createStmt, err := db.Prepare("INSERT INTO Categories (Name) VALUES (?)")

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[ERROR] Cannot prepare statement: %s", err.Error())
	}

	updateStmt, err := db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[ERROR] Cannot prepare statement: %s", err.Error())
	}

	categoryResult, err := tx.Stmt(createStmt).Exec(categoryName)

	// TODO: Remove
	tx.Rollback()
	return fmt.Errorf("[ERROR] Stop right here! %s", err.Error())

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[ERROR] Cannot create category: %s", err.Error())
	}

	fmt.Printf("Created category %q\n", categoryName)

	categoryID, _ := categoryResult.LastInsertId()
	updateStmtTx := tx.Stmt(updateStmt)
	var updatedRows int64

	for _, productID := range productIDs {
		updateResult, err := updateStmtTx.Exec(categoryID, productID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("[ERROR] Cannot update product #%d: %s", productID, err.Error())
		}

		count, err := updateResult.RowsAffected()

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("[ERROR] Cannot update product #%d: %s", productID, err.Error())
		}

		updatedRows += count
	}

	tx.Commit()
	fmt.Printf("Updated %d products\n", updatedRows)
	return nil
}

func transactionExample() {
	db, err := connectToDatabase()

	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connected to database")
	err = createAndUseCategory(db, "Miscellaneous", 1, 2)

	if err != nil {
		panic(err)
	}
}
