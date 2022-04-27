package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		// Recv() receives data from a channel
		// 1. Receive from channel
		// 2. if ok == true the channel is not closed
		// 3. walk() on received channel
		// 4. Keep receiving while ok == true (channel is open)
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
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
