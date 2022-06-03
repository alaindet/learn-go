package main

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	eatingTime = 0 * time.Second
	waitingTime = 0 * time.Second

	main()

	if len(orderFinished) != 5 {
		t.Errorf("wrong number of entries in slice")
	}
}
