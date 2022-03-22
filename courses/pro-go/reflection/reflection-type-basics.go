package main

import (
	"fmt"
	"reflect"
)

func getTypePath(t reflect.Type) string {
	path := t.PkgPath()
	if path == "" {
		return "built-in"
	}
	return path
}

func printReflectionTypesDetails(items ...interface{}) {
	for _, item := range items {
		itemType := reflect.TypeOf(item)
		fmt.Printf(
			"[Reflection Type] Name: %v, PkgPath: %v, Kind: %v\n",
			itemType.Name(),
			getTypePath(itemType),
			itemType.Kind(),
		)
	}
}

func reflectionTypeBasicsExample() {
	product := Product{"Kayak", "Watersports", 279.00}
	customer := Customer{"Alice", "New York"}
	payment := Payment{"USD", 100.50}

	printReflectionTypesDetails(
		product,
		customer,
		payment,
		123.4,
	)
}
