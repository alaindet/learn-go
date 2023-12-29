package main

import "fmt"

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

func isWorkingDay(d Weekday) bool {
	return !(d.Is(WeekdaySaturday) || d.Is(WeekdaySunday))
}
