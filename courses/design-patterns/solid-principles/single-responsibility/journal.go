package main

type Journal struct {
	entries      []string
	entriesCount int
}

func NewJournal() *Journal {
	return &Journal{}
}

func (j *Journal) AddEntry(entry string) int {
	j.entriesCount++
	if j.entries == nil {
		j.entries = make([]string, 0)
	}
	j.entries = append(j.entries, entry)
	return j.entriesCount
}

func (j *Journal) RemoveEntry(index int) int {
	j.entriesCount--

	if j.entriesCount < 0 {
		j.entriesCount = 0
	}

	entries := make([]string, 0, len(j.entries)-1)

	for i, entry := range j.entries {
		if i != index {
			entries = append(entries, entry)
		}
	}

	j.entries = entries

	return j.entriesCount
}
