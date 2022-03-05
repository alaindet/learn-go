package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name, Category string
	Price          float64
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

// This gets populated via init()
var Products = make(ProductData)

func init() {
	fmt.Println("product.go => init()")
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}
