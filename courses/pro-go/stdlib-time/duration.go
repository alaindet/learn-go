package main

import (
	"fmt"
	"strings"
	"time"
)

func timeDurationsBasics() {
	var d time.Duration = time.Hour + (30 * time.Minute)
	r := d.Round(time.Hour)    // 2 hours
	t := d.Truncate(time.Hour) // 1 hour

	fmt.Printf("Hours: %v\n", d.Hours())              // Hours: 1.5
	fmt.Printf("Mins: %v\n", d.Minutes())             // Mins: 90
	fmt.Printf("Seconds: %v\n", d.Seconds())          // Seconds: 5400
	fmt.Printf("Millseconds: %v\n", d.Milliseconds()) // Millseconds: 5400000

	fmt.Println(strings.Repeat("-", 10))
	fmt.Printf("Rounded Hours: %v\n", r.Hours())  // Rounded Hours: 2
	fmt.Printf("Rounded Mins: %v\n", r.Minutes()) // Rounded Mins: 120

	fmt.Println(strings.Repeat("-", 10))
	fmt.Printf("Truncated Hours: %v\n", t.Hours())  // Truncated Hours: 1
	fmt.Printf("Truncated Mins: %v\n", t.Minutes()) // Truncated Mins: 60
}

func changingTimeWithDurations() {
	t := time.Date(2022, 1, 1, 12, 1, 1, 1, time.UTC)
	diff := time.Duration(time.Hour)

	tAfter := t.Add(diff)
	tBefore := t.Add(-1 * diff)

	fmt.Println(t.Format(time.RFC3339))       // 2022-01-01T12:01:01Z
	fmt.Println(tAfter.Format(time.RFC3339))  // 2022-01-01T13:01:01Z
	fmt.Println(tBefore.Format(time.RFC3339)) // 2022-01-01T11:01:01Z

	// These two differ for a minute
	t1 := time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	t2 := t1.Add(time.Minute)

	// We round them by the hour
	r1 := t1.Round(time.Hour)
	r2 := t2.Round(time.Hour)

	fmt.Println(t1.Format(time.RFC3339)) // 2022-01-01T01:01:01Z
	fmt.Println(t2.Format(time.RFC3339)) // 2022-01-01T01:02:01Z
	fmt.Println(t1 == t2)                // false
	fmt.Println(r1.Format(time.RFC3339)) // 2022-01-01T01:00:00Z
	fmt.Println(r2.Format(time.RFC3339)) // 2022-01-01T01:00:00Z
	fmt.Println(r1 == r2)                // true
}

func durationRelativeToPointInTime() {
	future := time.Date(2050, 0, 0, 0, 0, 0, 0, time.UTC)
	past := time.Date(1950, 0, 0, 0, 0, 0, 0, time.UTC)

	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}

	yearsFromNow := toYears(time.Until(future))
	yearsPassed := toYears(time.Since(past))

	fmt.Printf("2050 is %d years from now\n", yearsFromNow) // 2050 is 27 years from now
	fmt.Printf("1950 was %d years ago\n", yearsPassed)      // 1950 was 72 years ago
}

func durationFromStrings() {
	d, err := time.ParseDuration("1h30m")
	_ = err
	fmt.Println(d) // 1h30m0s

	d, err = time.ParseDuration("0h25m3s234ms76us123ns")
	_ = err
	fmt.Println(d) // 25m3.234076123s

	// This does not work (no spaces allowed)!
	d, err = time.ParseDuration("25m 3s 234ms 76us 123ns")
	_ = err
	fmt.Println(d) // 0s
}
