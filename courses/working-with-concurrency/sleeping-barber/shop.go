package main

import (
	"time"

	"github.com/fatih/color"
)

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

func (s *BarberShop) AddBarber(barber string) {
	s.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(s.ClientsChan) == 0 {
				isSleeping = true
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
			}

			client, shopOpen := <-s.ClientsChan

			if !shopOpen {
				s.sendBarberHome(barber)
				return
			}

			if isSleeping {
				color.Yellow("%s wakes %s up.", client, barber)
				isSleeping = false
			}

			s.cutHair(barber, client)
		}
	}()
}

func (s *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(s.HairCutDuration)
	color.Green("%s is done cutting %s's hair.", barber, client)
}

func (s *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	s.BarbersDoneChan <- true
}

func (s *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")
	close(s.ClientsChan)
	s.Open = false

	for i := 0; i < s.NumberOfBarbers; i++ {
		<-s.BarbersDoneChan
	}

	close(s.BarbersDoneChan)

	color.Green("The barber shop is now closed, everyone is going home.")
}

func (s *BarberShop) AddClient(client string) {
	color.Green("*** %s arrives!", client)

	if !s.Open {
		color.Red("The shop is already closed, so %s leaves.", client)
		return
	}

	select {
	case s.ClientsChan <- client:
		color.Blue("%s takes a seat in the waiting room.", client)
	default:
		color.Red("The waiting room is full, so %s leaves.", client)
	}
}
