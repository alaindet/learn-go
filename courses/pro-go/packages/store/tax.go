package store

import "fmt"

const defaultTaxRate float64 = 0.2
const minThreshold = 10

// This is performed during initialization
var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

type taxRate struct {
	rate, threshold float64
}

// This is run when the package is loaded
// This is usually the first function to be declared
func init() {
	fmt.Println("tax.go => init()")
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
}

func newTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}
	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) (price float64) {

	price = product.price

	// Apply tax rate?
	if product.price > taxRate.threshold {
		price += (product.price * taxRate.rate)
	}

	// Cap the prize to its maximum allowed price?
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}

	return
}
