package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	onInit()
	registerGlobalFunctions()
	<-c
}

func onInit() {
	// ...
}

func registerGlobalFunctions() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	// ...
}

func add(this js.Value, args []js.Value) any {
	a, b, res := args[0].String(), args[1].String(), args[2].String()
	setInputValue(res, getInputValue(a)+getInputValue(b))
	return nil
}

func subtract(this js.Value, args []js.Value) any {
	a, b, res := args[0].String(), args[1].String(), args[2].String()
	setInputValue(res, getInputValue(a)-getInputValue(b))
	return nil
}

func getInputValue(id string) int {
	doc := js.Global().Get("document")
	rawValue := doc.Call("getElementById", id).Get("value").String()
	data, _ := strconv.Atoi(rawValue)
	return data
}

func setInputValue(id string, val int) {
	doc := js.Global().Get("document")
	doc.Call("getElementById", id).Set("value", val)
}
