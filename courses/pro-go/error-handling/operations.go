package main

import (
	"fmt"
)

// type CategoryError struct {
// 	requestedCategory string
// }

// func (err *CategoryError) Error() string {
// 	return fmt.Sprintf(
// 		"ERROR: Category %q does not exist",
// 		err.requestedCategory,
// 	)
// }

type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice(category string) (
	total float64,
	err error,
) {
	productCount := 0

	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}

	if productCount == 0 {
		// // Alternative
		// errorMessage := fmt.Sprintf("ERROR: Category %q does not exist", category)
		// err = errors.New(errorMessage)
		err = fmt.Errorf("ERROR: Category %q not find", category)
	}

	return
}

// Async wrapper for TotalPrice()
func (slice ProductSlice) TotalPriceAsync(
	categories []string,
	channel chan<- ChannelMessage,
) {
	for _, c := range categories {

		total, err := slice.TotalPrice(c)

		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}

	close(channel)
}
