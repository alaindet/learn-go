package main

import (
	"fmt"
	"syscall/js"
)

const (
	SayHelloFn = "sayHello"
)

// The empty channel just blocks the program
// so that JavaScript can use its functions
func main() {
	ch := make(chan struct{}, 0)
	fmt.Printf("Hello from Go WASM!\n") // You will see this in the console
	registerGlobalFunctions()
	<-ch
}

func registerGlobalFunctions() {
	js.Global().Set(SayHelloFn, SayHello())
}

func SayHello() js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			return "<p>Hello from GO WASM!</p>"
		},
	)
}
