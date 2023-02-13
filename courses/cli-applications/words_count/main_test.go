package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("the quick brown fox\n")
	expected := 4
	result := count(input, wordsMode)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("the\nquick\nbrown\nfox")
	expected := 4
	result := count(input, linesMode)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	input := bytes.NewBufferString("The quick brown fox jumps over the lazy dog")
	expected := 43
	result := count(input, bytesMode)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}
