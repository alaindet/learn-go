package main

import "time"

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func NewBarberShop(capacity int, d time.Duration) *BarberShop {
	return &BarberShop{
		ShopCapacity:    capacity,
		HairCutDuration: d,
		NumberOfBarbers: 0,
		BarbersDoneChan: make(chan bool),
		ClientsChan:     make(chan string, capacity),
		Open:            true,
	}
}

func (s *BarberShop) AddBarber(name string) {
	s.NumberOfBarbers++
	// TODO...
}
