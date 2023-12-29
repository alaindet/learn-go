package main

import (
	"fmt"
	"sync"
)

const (
	routines = 8
)

func main() {
	min := 100
	max := 1000
	delta := (max - min) / routines

	step1Ch := make(chan []int, routines)
	step2Ch := make(chan []int, routines)
	step3Ch := make(chan []int)

	go step1(step1Ch, delta, min, max)
	go step2(step1Ch, step2Ch)
	go step3(step2Ch, step3Ch, min, max)

	result := <-step3Ch
	count, lowest, highest := result[0], result[1], result[2]
	fmt.Println("count", count, "lowest", lowest, "highest", highest)
}

func step1(ch chan []int, delta, min, max int) {
	inf := min
	for i := 0; i < routines; i++ {
		if i > 0 {
			inf += delta + 1
		}

		sup := inf + delta

		if sup > max {
			sup = max
		}

		ch <- []int{inf, sup}
	}
	close(ch)
}

func step2(fromCh, toCh chan []int) {
	var wg sync.WaitGroup
	wg.Add(routines)
	for p := range fromCh {
		go func(p []int) {
			defer wg.Done()
			inf, sup := p[0], p[1]
			multiples := make([]int, 0, sup-inf)
			for i := inf; i <= sup; i++ {
				if isMultipleOf(i, 39) {
					multiples = append(multiples, i)
				}
			}
			toCh <- multiples
		}(p)
	}
	wg.Wait()
	close(toCh)
}

func step3(fromCh, toCh chan []int, min, max int) {
	count := 0
	lowest := max
	highest := min
	for proc := range fromCh {
		count += len(proc)
		for _, n := range proc {
			if n < lowest {
				lowest = n
			}
			if n > highest {
				highest = n
			}
		}
	}
	toCh <- []int{count, lowest, highest}
}

func isMultipleOf(n, m int) bool {
	return n%m == 0
}
