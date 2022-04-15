package main

import (
	"fmt"
	"mymodule/foo"
	"mymodule/foo/bar"
)

func main() {
	fmt.Println("mymodule")
	foo.FooFn()
	bar.BarFn()
}
