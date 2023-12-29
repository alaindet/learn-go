package main

import "fmt"

func iotaExample() {

	seasons := []Season{
		SeasonWinter,
		SeasonSpring,
		SeasonSummer,
		SeasonAutumn,
		42,
	}

	for _, s := range seasons {

		if !s.Allowed() {
			fmt.Printf("Season %s does not exist\n", s)
			continue
		}

		t := getSeasonTemperature(s)
		fmt.Printf("Season %s has a %s temperature\n", s, t)
	}
}

func getSeasonTemperature(s Season) string {
	switch s {
	case SeasonWinter:
		return "cold"
	case SeasonSummer:
		return "hot"
	case SeasonSpring:
		fallthrough
	case SeasonAutumn:
		return "moderate"
	default:
		return ""
	}
}
