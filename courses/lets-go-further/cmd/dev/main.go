package main

import (
	"fmt"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    int     `db:"id" json:"id"`
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

func main() {
	// 1. Connect to SQLite file (creates app.db automatically if it doesn't exist)
	db, err := sqlx.Connect("sqlite", "app.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to local SQLite database file!")

	// 2. Create Schema
	schema := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
	);`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	// 3. Insert Data (using NamedExec)
	insertQuery := `INSERT INTO products (name, price) VALUES (:name, :price)`
	newProduct := Product{Name: "Some Cool Product", Price: 99.99}

	_, err = db.NamedExec(insertQuery, newProduct)
	if err != nil {
		log.Fatalf("Failed to insert product: %v", err)
	}
	fmt.Println("Inserted product:", newProduct.Name)

	// 4. Query Data (using sqlx.Select to map rows into structs)
	var products []Product
	err = db.Select(&products, "SELECT id, name, price FROM products")
	if err != nil {
		log.Fatalf("Failed to query products: %v", err)
	}

	fmt.Println("\n--- All Products ---")
	for _, p := range products {
		fmt.Printf("[%d] %s - $%.2f\n", p.ID, p.Name, p.Price)
	}
}
