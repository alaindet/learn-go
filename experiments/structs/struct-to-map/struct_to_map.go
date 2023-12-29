package main

import (
	"fmt"
	"reflect"
)

// This converts most structs into maps
// Supports nested structs
// Supported types mimics JSON format, so only numbers, booleans, arrays/slices
// and nested structs are supported
// Functions, channels, interfaces and pointers are skipped
func StructToMap(input any) map[string]any {

	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	return structToMap(t, v)
}

func structToMap(t reflect.Type, v reflect.Value) map[string]any {

	if t.Kind() != reflect.Struct {
		return make(map[string]any)
	}

	fieldsCount := v.NumField()
	result := make(map[string]any, fieldsCount)

	for i := 0; i < fieldsCount; i++ {
		f := t.Field(i)
		tt := f.Type
		vv := v.Field(i)
		item, err := structFieldToMapItem(tt, vv)
		if err == nil {
			result[f.Name] = item
		}
	}

	return result
}

func structFieldToMapItem(t reflect.Type, v reflect.Value) (any, error) {

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	switch t.Kind() {
	case reflect.Struct:
		return structToMap(t, v), nil
	case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return nil, fmt.Errorf("unsupported type %v", t)
	case reflect.Array, reflect.Slice:
		count := v.Len()
		list := make([]any, 0, count)

		for i := 0; i < count; i++ {
			tt := t.Elem()
			vv := v.Index(i)
			item, err := structFieldToMapItem(tt, vv)
			if err == nil {
				list = append(list, item)
			}
		}

		return list, nil
	default:
		return v, nil
	}
}
