package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	valuesCount := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		valuesCount = val.Len()
		getField = val.Index
	case reflect.Struct:
		valuesCount = val.NumField()
		getField = val.Field
	}

	for i := 0; i < valuesCount; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {

	// Extracts the value of the concrete value passed to x,
	// since x is actually an interface implemented by the concrete value
	val := reflect.ValueOf(x)

	// Example
	// typ := reflect.TypeOf(x)
	// fmt.Println("concrete type", typ) // struct { Name string }

	// Note: reflect.Pointer is equivalent to reflect.Ptr
	if val.Kind() == reflect.Pointer {
		val = val.Elem() // For pointers, it dereferences them
	}

	return val
}
