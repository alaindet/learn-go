package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(pos int) bool {
	return pos >= t.Start && pos <= t.End
}

type GoodFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewGoodFormattedText(plainText string) *GoodFormattedText {
	return &GoodFormattedText{plainText, nil}
}

func (f *GoodFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	f.formatting = append(f.formatting, r)
	return r
}

func (f *GoodFormattedText) String() string {
	b := strings.Builder{}

	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		for _, r := range f.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		b.WriteRune(rune(c))
	}

	return b.String()
}

func goodFormatterExample() {
	t := "This is a brave new word"
	f := NewGoodFormattedText(t)
	f.Range(10, 15).Capitalize = true
	fmt.Println(f.String())
}
