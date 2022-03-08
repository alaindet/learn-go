package main

import "fmt"

func p(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func scanFromStandardInput() {
	var name string
	var category string
	var price float64

	fmt.Print("Enter text to scan: ")

	// Basic: accepts anything then splits by spaces, newlines are like spaces
	// n, err := fmt.Scan(&name, &category, &price)

	// Like Scan(), but breaks on newlines
	// n, err := fmt.Scanln(&name, &category, &price)

	// Like Scan(), but input is a value, not standard input
	// source := "Lifejacket Watersports 48.95">
	// n, err := fmt.Sscan(source, &name, &category, &price)

	// Scan a value via a template string
	// Similar to a regex but much simpler, it only splits by space
	// source := "Values are: Lifejacket Watersports 48.95"
	// template := "Values are: %s %s %f"
	// n, err := fmt.Sscanf(source, template, &name, &category, &price)

	// Scan standard input via a template string
	template := "Values are: %s %s %f"
	n, err := fmt.Scanf(template, &name, &category, &price)

	if err == nil {
		p("Scanned %v values", n)
		p("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		p("Error: %v", err.Error())
	}
}

// Args of same value
func scanIntoSlice() {
	vals := make([]string, 3)
	ivals := make([]interface{}, 3)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}

	fmt.Print("Enter text to scan: ")
	fmt.Scan(ivals...)
	p("Name: %v", vals)
}
