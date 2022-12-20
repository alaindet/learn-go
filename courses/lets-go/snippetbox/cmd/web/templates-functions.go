package main

import (
	"html/template"
	"reflect"
	"time"
)

var templateFunctions = template.FuncMap{
	"friendlyDate": friendlyDate,
	"isLastItem":   isLastItem,
}

func friendlyDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func isLastItem(index int, list interface{}) bool {
	return index == reflect.ValueOf(list).Len()-1
}
