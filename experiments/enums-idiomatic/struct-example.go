package main

import "fmt"

type PaymentStatus struct {
	code    int
	message string
}

func (s PaymentStatus) String() string {
	return s.message
}

func (s PaymentStatus) Equals(other PaymentStatus) bool {
	return s.code == other.code && s.message == other.message
}

var (
	PaymentStatusPending  = PaymentStatus{code: 0, message: "Payment is pending"}
	PaymentStatusApproved = PaymentStatus{code: 1, message: "Payment approved"}
	PaymentStatusDeclined = PaymentStatus{code: 2, message: "Payment declined"}
)

/**
 * Pros
 * - Each instance can be as complex as needed (fields, methods)
 *
 * Cons
 * - Run-time values (not constants)
 * - Verbose to setup
 * - Slower than iota-based and string-based enums
 * - Needs marshalling/unmarshalling for JSON serialization
 */
func structEnumExample() {
	fmt.Println("\n# Struct enum example")

	myPaymentStatus := PaymentStatusApproved
	myCustomPayment := PaymentStatus{code: 1, message: "Payment approved"}

	// Mind this: since PaymentStatus is a struct with only comparable fields
	// it is itself comparable
	//
	// Here, the switch compares fields like the custom Equals method defined
	// above
	switch myPaymentStatus {
	case PaymentStatusPending:
		fmt.Println("Pending")
	case PaymentStatusApproved:
		fmt.Println("Approved")
	case PaymentStatusDeclined:
		fmt.Println("Declined")
	}

	fmt.Printf("Same stati? %t\n", myPaymentStatus.Equals(myCustomPayment))
}
