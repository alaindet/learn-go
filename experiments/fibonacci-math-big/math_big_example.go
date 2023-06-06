package main

import (
	"fmt"
	"math/big"
	"time"
)

func mathBigExample() {
	calculate := make(chan struct{})
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			calculate <- struct{}{}
		}
		quit <- struct{}{}
	}()

	fibonacciMathBig(calculate, quit)
}

func fibonacciMathBig(calculate, quit chan struct{}) {

	xOld := big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(1)

	for {
		select {
		case <-calculate:
			xOld.Set(x)
			x.Set(y)
			y.Add(xOld, y)
			fmt.Println(y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
