package main

import (
	"fmt"
	"regexp"
	"strings"
)

func regexBasicMatching() {
	pattern := "[A-z]oat" // [A-z] matches all letters, lower and upper case
	s := "A boat for one person"

	if match, err := regexp.MatchString(pattern, s); err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	regex, err := regexp.Compile(pattern)
	fmt.Printf("%T\n", regex) // *regexp.Regexp
	s2 := "Is that a goat?"
	s3 := "I like oats"

	if err != nil {
		fmt.Println("Compile error", err)
	} else {
		fmt.Println(
			regex.MatchString(s),  // true
			regex.MatchString(s2), // true
			regex.MatchString(s3), // false
		)
	}

	// Common methods of Regexp
	// MatchString(s)
	// FindStringIndex(s)
	// FindAllStringIndex(s, max)
	// FindString(s)
	// FindAllString(s, max)
	// Split(s, max)

	fmt.Println(
		regex.FindString("I have a coat"),          // "coat"
		regex.FindString("I have an apricot"),      // ""
		regex.FindStringIndex("I have a coat"),     // [9 13]
		regex.FindStringIndex("I have an apricot"), // []
	)
}

func regexCompiledPattern() {
	printMatch := func(
		label string,
		index int,
		s string,
		stringIndices []int,
	) string {

		if label == "" {
			label = "printMatch"
		}

		return fmt.Sprintf(
			"%s: #%d => [%d, %d] => %s",
			label,
			index,
			stringIndices[0],
			stringIndices[1],
			string(s[stringIndices[0]:stringIndices[1]]),
		)
	}

	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."

	firstIndex := pattern.FindStringIndex(description)
	fmt.Println(printMatch("FindStringIndex", 0, description, firstIndex))

	allIndices := pattern.FindAllStringIndex(description, -1)
	for i, indices := range allIndices {
		fmt.Println(printMatch("FindAllStringIndex", i, description, indices))
	}
}

func regexSplit() {
	description := "Kayak. A boat for one person."
	pattern := regexp.MustCompile(" |boat|one")
	split := pattern.Split(description, -1)
	joined := strings.Join(split, "§")
	fmt.Println(joined) // Kayak.§A§§§for§§§person.
}

func regexSubexpressions() {
	description := "Kayak. A boat for one person."
	pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	matched := pattern.FindStringSubmatch(description) // <-- Look here
	joined := strings.Join(matched, "§")
	fmt.Println(joined) // A boat for one person§boat§one
}
