package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		// From docs: Helper marks the calling function as a test helper function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// This is a subtest
	t.Run(
		"saying hello to people",
		func(t *testing.T) {
			got := Hello("Foo", "")
			want := "Hello, Foo"
			assertCorrectMessage(t, got, want)
		},
	)

	t.Run(
		"say \"Hello, World\" when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		},
	)

	t.Run(
		"in Italian",
		func(t *testing.T) {
			got := Hello("Bar", "it")
			want := "Ciao, Bar"
			assertCorrectMessage(t, got, want)
		},
	)
}
