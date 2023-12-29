package main

import (
	"encoding/json"
	"fmt"
)

func enumerExample() {

	// Convert a constant to string
	fmt.Println(MonthAugust.String())

	// Use constant as string automatically
	fmt.Printf("Month is: %q\n", MonthMay)

	// Convert string to constant int
	// TODO: Remove prefix?
	m, err := MonthString("MonthMarch")
	if err != nil {
		fmt.Println("Unknown month", err)
	} else {
		fmt.Printf("Known month: %q\n", m)
	}

	for _, m := range MonthValues() {
		fmt.Println("Listing months", m)
	}

	var invalidMonth Month = 15
	fmt.Printf("Is %d a valid month? %t\n", invalidMonth, invalidMonth.IsAMonth())

	var validMonth Month = 6
	fmt.Printf("Is %d a valid month? %t\n", validMonth, validMonth.IsAMonth())

	monthJSON, err := MonthSeptember.MarshalJSON()
	if err != nil {
		fmt.Printf("Could not convert %q to JSON", MonthSeptember)
	} else {
		fmt.Println(string(monthJSON))
	}

	myMap := map[string]Month{
		"jan": MonthJanuary,
		"apr": MonthApril,
	}

	myMapJSON, _ := json.Marshal(myMap)
	fmt.Println(string(myMapJSON))
	// Prints {"apr":"MonthApril","jan":"MonthJanuary"}
}
