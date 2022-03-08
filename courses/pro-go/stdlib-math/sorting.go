package main

import (
	"sort"
)

func sortingBasics() {

	// Floats
	floats := []float64{1.23, 4.56, 1.01}
	p("Are %v sorted? %t", floats, sort.Float64sAreSorted(floats))
	sort.Float64s(floats)
	p("Are %v sorted? %t", floats, sort.Float64sAreSorted(floats))

	// Integers
	ints := []int{3, 2, 7, 1}
	p("Are %v sorted? %t", ints, sort.IntsAreSorted(ints))
	sort.Ints(ints)
	p("Are %v sorted? %t", ints, sort.IntsAreSorted(ints))

	// Strings
	strs := []string{"bb", "cc", "aa"}
	p("Are %v sorted? %t", strs, sort.StringsAreSorted(strs))
	sort.Strings(strs)
	p("Are %v sorted? %t", strs, sort.StringsAreSorted(strs))
}

func immutableIntSort(ints []int) []int {
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	return sortedInts
}

func searchValuesInSlices() {
	ints := []int{5, 4, 2, 1, 9}
	p("%v %p", ints, &ints) // [5 4 2 1 9] 0xc00000c030
	sortedInts := immutableIntSort(ints)
	p("%v %p", sortedInts, &sortedInts) // [1 2 4 5 9] 0xc00000c060

	indexOf5 := sort.SearchInts(sortedInts, 5)
	p("Index of 5: %d", indexOf5) // [1 2 4 5 9] => 3
}
