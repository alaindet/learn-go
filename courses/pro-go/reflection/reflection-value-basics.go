package main

import (
	"fmt"
	"reflect"
)

var INT_PTR_TYPE = reflect.TypeOf((*int)(nil))
var BYTE_SLICE_TYPE = reflect.TypeOf([]byte(nil))

func printReflectionValueDetails(items ...interface{}) {
	for _, item := range items {
		itemType := reflect.TypeOf(item)
		itemValue := reflect.ValueOf(item)

		if itemType == INT_PTR_TYPE {
			fmt.Printf("*Int: %v\n", itemValue.Elem().Int())
			continue
		}

		if itemType == BYTE_SLICE_TYPE {
			fmt.Printf("Byte slice: %v\n", itemValue.Bytes())
			continue
		}

		switch itemValue.Kind() {
		case reflect.Bool:
			var val bool = itemValue.Bool()
			fmt.Printf("Bool: %v\n", val)
		case reflect.Int:
			var val int64 = itemValue.Int()
			fmt.Printf("Int: %v\n", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = itemValue.Float()
			fmt.Printf("Float: %v\n", val)
		case reflect.String:
			var val string = itemValue.String()
			fmt.Printf("String: %v\n", val)
		case reflect.Ptr:
			var val reflect.Value = itemValue.Elem()
			if val.Kind() == reflect.Int {
				fmt.Printf("Pointer to Int: %v\n", val.Int())
			}
		default:
			fmt.Printf("Other: %v\n", itemValue.String())
		}
	}
}

func reflectionValueBasicsExample() {
	product := Product{"Kayak", "Watersports", 279.00}
	aNumber := 100
	names := []byte("Alice")

	printReflectionValueDetails(
		true,
		10,
		2.34,
		"Alice",
		product,
		&aNumber,
		names,
	)
}
