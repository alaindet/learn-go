package main

import "testing"

func assertStrings(t testing.TB, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func assertError(t testing.TB, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("Result: %q Expected: %q", result, expected)
	}
}

func TestSearch(t *testing.T) {

	d := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		result, err := d.Search("test")
		_ = err
		expected := "this is just a test"
		assertStrings(t, result, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		result, err := d.Search("unknown_word")
		_ = result
		expected := ErrWordNotFound
		assertError(t, err, expected)
	})
}

func TestAdd(t *testing.T) {

	// WARNING:
	// Do this --> var d = map[string]string{}
	// Or this --> var d = make(map[string]string)
	// BUT NOT THIS --> var d map[string]string
	// Because the last one is a nil map to which you cannot add values
	d := Dictionary{}

	t.Run("add definition", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d.Add(word, def)
		expected := def
		result, err := d.Search(word)

		if err != nil {
			t.Fatal("should find added word", err)
		}

		if result != expected {
			t.Errorf("Result: %q Expected: %q", result, expected)
		}
	})
}
