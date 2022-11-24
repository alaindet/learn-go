package main

import (
	"fmt"
	"strings"
	"unicode"
)

type BadFormattedText struct {
	plainText  string
	capitalize []bool
}

func NewBadFormattedText(plainText string) *BadFormattedText {
	return &BadFormattedText{plainText, make([]bool, len(plainText))}
}

func (f *BadFormattedText) String() string {
	b := strings.Builder{}

	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			b.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			b.WriteRune(rune(c))
		}
	}

	return b.String()
}

func (f *BadFormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func badFormatterExample() {
	t := "This is a brave new word"
	f := NewBadFormattedText(t)
	f.Capitalize(10, 15)
	fmt.Println(f.String())
}
