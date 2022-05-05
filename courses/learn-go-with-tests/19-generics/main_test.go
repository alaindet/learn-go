package main

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	// This gives a compilation error with generics!
	// t.Run("compare apples and oranges", func(t *testing.T) {
	// 	AssertEqual(t, "1", 1)
	// })
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")

		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, result, expected T) {
	t.Helper()
	if result != expected {
		t.Errorf("got %+v, want %+v", result, expected)
	}
}

func AssertNotEqual[T comparable](t *testing.T, result, expected T) {
	t.Helper()
	if result == expected {
		t.Errorf("didn't want %+v", result)
	}
}

func AssertTrue(t *testing.T, result bool) {
	t.Helper()
	if !result {
		t.Errorf("got %v, want true", result)
	}
}

func AssertFalse(t *testing.T, result bool) {
	t.Helper()
	if result {
		t.Errorf("got %v, want false", result)
	}
}
