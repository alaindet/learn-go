package main

import "fmt"

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) int {
	entryCount--
	if entryCount < 0 {
		entryCount = 0
	}

	// TODO...

	return entryCount
}

func main() {

}
