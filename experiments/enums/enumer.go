package main

type Month int

//go:generate go run github.com/dmarkham/enumer -type=Month -output enumer_enum.go -json -yaml
const (
	MonthNone Month = iota
	MonthJanuary
	MonthFebruary
	MonthMarch
	MonthApril
	MonthMay
	MonthJune
	MonthJuly
	MonthAugust
	MonthSeptember
	MonthOctober
	MonthNovember
	MonthDecember
)
