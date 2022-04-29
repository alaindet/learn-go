package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Arabic int
	Roman  string
}

var cases []TestCase = []TestCase{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		name := fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman)
		t.Run(name, func(t *testing.T) {
			result := ConvertToRoman(test.Arabic)
			expected := test.Roman
			if result != expected {
				t.Errorf("result: %q, expected: %q", result, expected)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:1] {
		name := fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic)
		t.Run(name, func(t *testing.T) {
			result := ConvertToArabic(test.Roman)
			expected := test.Arabic
			if result != expected {
				t.Errorf("result: %q, expected: %q", result, expected)
			}
		})
	}
}
