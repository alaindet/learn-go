package main

import "slices"

type Season int

const (
	SeasonUnknown Season = iota
	SeasonWinter
	SeasonSpring
	SeasonSummer
	SeasonAutumn
)

func (s Season) String() string {
	switch s {
	case SeasonUnknown:
		return "SeasonUnknown"
	case SeasonWinter:
		return "SeasonWinter"
	case SeasonSpring:
		return "SeasonSpring"
	case SeasonSummer:
		return "SeasonSummer"
	case SeasonAutumn:
		return "SeasonAutumn"
	default:
		return ""
	}
}

func (s Season) Is(other Season) bool {
	return s.String() == other.String()
}

func (s Season) Allowed() bool {

	allowedList := []Season{
		SeasonUnknown,
		SeasonWinter,
		SeasonSpring,
		SeasonSummer,
		SeasonAutumn,
	}

	return slices.Contains(allowedList, s)
}
