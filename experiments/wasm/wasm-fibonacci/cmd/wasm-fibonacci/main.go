package main

import (
	"syscall/js"

	"wasm_fibonacci/cmd/fibonacci"
)

func main() {
	c := make(chan struct{}, 0)
	registerGlobalFunctions()
	<-c
}

func registerGlobalFunctions() {
	js.Global().Set("fibonacci", js.FuncOf(wasmFibonacci))
}

func wasmFibonacci(this js.Value, args []js.Value) any {
	n := args[0].Int()
	return fibonacci.Fibonacci(n)
}
