package main

import "fmt"

func p(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
