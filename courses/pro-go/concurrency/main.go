package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	time.Sleep(time.Second * 2)
	fmt.Println("main function complete")
}
