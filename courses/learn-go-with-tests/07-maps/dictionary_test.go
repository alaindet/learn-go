package main

import "testing"

func assertDefinition(
	t testing.TB,
	dictionary Dictionary,
	word, definition string,
) {
	t.Helper()

	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != result {
		t.Errorf("Result: %q Expected:  %q", result, definition)
	}
}

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

	t.Run("add new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d.Add(word, def)
		assertDefinition(t, d, word, def)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{word: def}
		err := d.Add(word, "new definition")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{word: def}
		newDef := "new definition"
		err := d.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, d, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		d := Dictionary{}
		err := d.Update(word, def)
		assertError(t, err, ErrWordNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "this is ignored"}
		d.Delete(word)
		_, err := d.Search(word)
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{}
		err := d.Delete(word)
		assertError(t, err, ErrWordNotFound)
	})
}
