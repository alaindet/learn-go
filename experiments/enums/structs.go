package main

import "slices"

type Weekday struct {
	name string
}

func (w Weekday) String() string {
	return w.name
}

func (w Weekday) Is(other Weekday) bool {
	return w.name == other.name
}

func (w Weekday) Allowed() bool {

	allowedList := []string{
		"monday",
		"tuesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}

	return slices.Contains(allowedList, w.name)
}

var (
	WeekdayUnknown  = Weekday{""}
	WeekdayMonday   = Weekday{"monday"}
	WeekdayTuesday  = Weekday{"tuesday"}
	WeekdayThursday = Weekday{"thursday"}
	WeekdayFriday   = Weekday{"friday"}
	WeekdaySaturday = Weekday{"saturday"}
	WeekdaySunday   = Weekday{"sunday"}
)
