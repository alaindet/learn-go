package main

import "fmt"

type Status string

const (
	StatusUnknown Status = ""
	StatusPending Status = "pending"
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

/**
 * Pros
 * - Compile-time constants
 * - Type-safe
 * - Easy to write
 * - No String() method needed as string is the underlying type
 * - Slower than iota integers, but fast enough
 *
 * Cons
 * - Instances have the enum's name prefixed
 */
func stringEnumExample() {
	fmt.Println("\n# String enum example")

	myStatus := StatusSuccess
	myOtherStatus := "success"

	// This does not compile as they are of different types
	// fmt.Println("Same stati?", myStatus == myOtherStatus)

	// This works as the type is converted first
	myThirdStatus := Status(myOtherStatus)
	fmt.Printf("Same stati? %t\n", myStatus == myThirdStatus)

	var unknownStatus Status
	fmt.Printf("Is status unknown? %t\n", unknownStatus == StatusUnknown)
}
