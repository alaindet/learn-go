package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	// Extracts the value of the concrete value passed to x,
	// since x is actually an interface implemented by the concrete value
	val := reflect.ValueOf(x)

	// Example
	// typ := reflect.TypeOf(x)
	// fmt.Println("concrete type", typ) // struct { Name string }

	// DANGER: Assuming val is a struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
