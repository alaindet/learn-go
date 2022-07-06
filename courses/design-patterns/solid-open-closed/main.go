package main

import "fmt"

/*
OCP - Open-closed principle
Any class (or struct) should be open for extension, but closed for modification
*/
func main() {
	products := []Product{
		{"Apple", green, small},
		{"Tree", green, large},
		{"House", blue, large},
	}

	f := Filter{}

	// Green products
	greenSpec := ColorSpecification{green}
	greenProducts := f.Filter(products, greenSpec)
	for _, p := range greenProducts {
		fmt.Printf("Green: %s\n", p.Name)
	}
	fmt.Printf("\n\n")

	// Large products
	largeSpec := SizeSpecification{large}
	largeProducts := f.Filter(products, largeSpec)
	for _, p := range largeProducts {
		fmt.Printf("Large: %s\n", p.Name)
	}
	fmt.Printf("\n\n")

	// Large and blue products
	blueSpec := ColorSpecification{blue}
	largeAndBlueSpc := AndSpecification{blueSpec, largeSpec}
	largeAndBlueProducts := f.Filter(products, largeAndBlueSpc)
	for _, p := range largeAndBlueProducts {
		fmt.Printf("Large and blue: %s\n", p.Name)
	}
	fmt.Printf("\n\n")
}
