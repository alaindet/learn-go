package main

import "fmt"

func main() {
	iotaExample()
	structExample()
}

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

func structExample() {

	days := []Weekday{
		WeekdayMonday,
		WeekdaySunday,
		Weekday{"nope"},
	}

	for _, d := range days {

		if !d.Allowed() {
			fmt.Printf("Working day %q does not exist\n", d)
			continue
		}

		working := isWorkingDay(d)
		fmt.Printf("Working on %s? %t\n", d, working)
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

func isWorkingDay(d Weekday) bool {
	return !(d.Is(WeekdaySaturday) || d.Is(WeekdaySunday))
}
