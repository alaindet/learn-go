package main

import (
	"fmt"
	"reflect"
	"strings"
)

func printDetails(items ...interface{}) {
	for _, item := range items {
		fieldDetails := []string{}
		itemType := reflect.TypeOf(item)
		itemValue := reflect.ValueOf(item)

		// Non-struct types
		if itemType.Kind() != reflect.Struct {
			fmt.Printf("[%v] %v\n", itemType.Name(), itemValue)
			continue
		}

		// Struct types
		for i := 0; i < itemType.NumField(); i++ {
			fieldName := itemType.Field(i).Name
			fieldValue := itemValue.Field(i)
			fieldDetails = append(
				fieldDetails,
				fmt.Sprintf("%v: %v", fieldName, fieldValue),
			)
		}

		fmt.Printf("[%v] %v\n", itemType.Name(), strings.Join(fieldDetails, ", "))
	}
}

func reflectionBasicsExample() {
	product := Product{Name: "Kayak", Category: "Watersports", Price: 279}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment{Currency: "USD", Amount: 100.50}

	printDetails(
		product,
		customer,
		payment,
	)
}
