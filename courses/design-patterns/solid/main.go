package main

import (
	"fmt"
)

func main() {
	j := NewJournal()
	j.AddEntry("Hello there, how are you?")
	j.AddEntry("Fine thanks. And you?")
	j.AddEntry("Fine too! Have a nice day")
	j.RemoveEntry(1)

	p := NewPersistence("\n")
	p.SaveToFile(j, "journal.txt")

	fmt.Printf("entries (%T): %v\n", j.entries, j.entries)
}
