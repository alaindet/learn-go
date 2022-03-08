package main

import (
	"fmt"
	"time"
)

func createDatesBasics() {
	// 2022-03-08 17:56:21.6403112 +0100 CET m=+0.000081001
	fmt.Println(time.Now())

	// 2022-03-03 12:12:12 +0100 CET
	fmt.Println(time.Date(2022, 3, 3, 12, 12, 12, 0, time.Local))

	// 1970-01-02 11:17:36 +0100 CET
	fmt.Println(time.Unix(123456, 0))
}

func timeInstanceBasics() {
	t := time.Now()
	fmt.Println(t.Date()) // 2022 March 8
}
