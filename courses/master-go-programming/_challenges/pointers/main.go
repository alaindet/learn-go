package main

import (
	"fmt"
)

func exercise1() {
	x := 10.10
	p := &x

	fmt.Println("x  - Value of x =", x)
	fmt.Println("&x - Address of x =", &x)
	fmt.Println("p  - Pointer to x =", p)
	fmt.Println("&p - Address of pointer to x =", &p)
	fmt.Println("*p - Value of pointed variable =", *p)
	// x  - Value of x = 10.1
	// &x - Address of x = 0xc000014068
	// p  - Pointer to x = 0xc000014068
	// &p - Address of pointer to x = 0xc00000e028
	// *p - Value of pointed variable = 10.1
}

func exercise2() {
	x, y := 10, 2
	xPtr, yPtr := &x, &y
	z := *xPtr / *yPtr
	fmt.Println(z) // 5
}

func swap(a *float64, b *float64) {
	*b, *a = *a, *b
}

func exercise3() {
	x, y := 5.5, 8.8
	swap(&x, &y)
	fmt.Println(x, y) // 8.8 5.5
}

func main() {
	exercise1()
	exercise2()
	exercise3()
}
