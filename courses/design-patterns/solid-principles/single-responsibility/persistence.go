package main

/*
The persistence is a separate concern than the journal itself, and it can also
be shared between different entitities
*/

import (
	"os"
	"strings"
)

type Persistence struct {
	lineSeparator string
}

func NewPersistence(lineSeparator string) *Persistence {
	return &Persistence{lineSeparator: lineSeparator}
}

func (p *Persistence) SaveToFile(j *Journal, filename string) error {
	content := []byte(strings.Join(j.entries, p.lineSeparator))

	err := os.WriteFile(filename, content, 0644)

	if err != nil {
		return err
	}

	return nil
}
