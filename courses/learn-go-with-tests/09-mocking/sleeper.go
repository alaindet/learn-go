package main

import "time"

// Strange, but this does not natively exist on "time" package
type Sleeper interface {
	Sleep()
}

// The testing sleeper
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// The real sleeper
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
