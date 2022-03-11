package main

import (
	"fmt"
	"strings"
	"time"
)

func createDatesBasics() {
	// 2022-03-08 17:56:21.6403112 +0100 CET m=+0.000081001
	fmt.Println(time.Now())

	// 2022-03-03 12:12:12 +0100 CET
	fmt.Println(time.Date(2022, 3, 3, 12, 12, 12, 0, time.Local))

	// 1970-01-02 11:17:36 +0100 CET
	fmt.Println(time.Unix(123456, 0))
}

func timeInstanceBasics() {
	t := time.Now()
	y, m, d := t.Date()
	h, i, s := t.Clock()

	// 2022 March 11 13 04 31
	fmt.Println(y, m, d, h, i, s)

	// int time.Month int int int int
	fmt.Printf("%T %T %T %T %T %T\n", y, m, d, h, i, s)

	fmt.Println(
		t.Year(),       // 2022 // <-- int
		t.YearDay(),    // 70 // <-- int
		t.Month(),      // March // <-- time.Month
		t.Day(),        // 11 // <-- int
		t.Weekday(),    // Friday // <-- time.Weekday
		t.Hour(),       // 16 // <-- int
		t.Minute(),     // 36 // <-- int
		t.Second(),     // 57 // <-- int
		t.Nanosecond(), // 905059700 // <-- int
	)

	fmt.Printf("%T\n", t.Nanosecond())
}

func formattingTime() {

	t := time.Now()

	// Custom layout
	// Reference time => https://pkg.go.dev/time#pkg-constants
	layout := "Day is: 02, Month is: 01, Year is: 2006"
	fmt.Println(t.Format(layout)) // Day is: 11, Month is: 03, Year is: 2022

	// Standard layouts
	fmt.Println(t.Format(time.ANSIC))       // Fri Mar 11 16:57:35 2022
	fmt.Println(t.Format(time.UnixDate))    // Fri Mar 11 16:57:35 CET 2022
	fmt.Println(t.Format(time.RubyDate))    // Fri Mar 11 16:57:35 +0100 2022
	fmt.Println(t.Format(time.RFC822))      // 11 Mar 22 16:57 CET
	fmt.Println(t.Format(time.RFC822Z))     // 11 Mar 22 16:57 +0100
	fmt.Println(t.Format(time.RFC850))      // Friday, 11-Mar-22 16:57:35 CET
	fmt.Println(t.Format(time.RFC1123))     // Fri, 11 Mar 2022 16:57:35 CET
	fmt.Println(t.Format(time.RFC1123Z))    // Fri, 11 Mar 2022 16:57:35 +0100
	fmt.Println(t.Format(time.RFC3339))     // 2022-03-11T16:57:35+01:00
	fmt.Println(t.Format(time.RFC3339Nano)) // 2022-03-11T16:57:35.500529+01:00
	fmt.Println(t.Format(time.Kitchen))     // 4:57PM
	fmt.Println(t.Format(time.Stamp))       // Mar 11 16:57:35
	fmt.Println(t.Format(time.StampMilli))  // Mar 11 16:57:35.500
	fmt.Println(t.Format(time.StampMicro))  // Mar 11 16:57:35.500529
	fmt.Println(t.Format(time.StampNano))   // Mar 11 16:57:35.500529000
}

func parsingTime() {
	datesToParse := []string{
		"1990-Aug-15",
		"2001-Sep-11",
	}

	customTimeLayout := "2006-Jan-02"
	for _, date := range datesToParse {
		parsedTime, err := time.Parse(customTimeLayout, date)
		if err != nil {
			fmt.Printf("Error %s\n", err.Error())
			continue
		}
		fmt.Printf("%q => %v\n", date, parsedTime)
	}
	// "1990-Aug-15" => 1990-08-15 00:00:00 +0000 UTC
	// "2001-Sep-11" => 2001-09-11 00:00:00 +0000 UTC

	datesToParse2 := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 MST",
	}

	for _, date := range datesToParse2 {
		parsedTime, err := time.Parse(time.RFC822, date)
		if err != nil {
			fmt.Printf("Error %s\n", err.Error())
			continue
		}
		fmt.Printf("%q => %v\n", date, parsedTime)
	}
	// "09 Jun 95 00:00 GMT" => 1995-06-09 00:00:00 +0000 GMT
	// "02 Jun 15 00:00 MST" => 2015-06-02 00:00:00 +0000 MST
}

func changingTime() {
	t := time.Now().UTC()

	t2 := t.AddDate(0, 0, 1)  // Add 1 day
	t3 := t.AddDate(0, 1, 0)  // Add 1 month
	t4 := t.AddDate(1, 0, 0)  // Add 1 year
	t5 := t.AddDate(0, 0, -1) // Substract 1 day

	fmt.Println(t2, t3, t4, t5)

	today := time.Now().UTC()
	todayAgain := time.Now().UTC()
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)
	fmt.Println(
		today.After(yesterday),  // true
		today.After(tomorrow),   // false
		today.Before(yesterday), // false
		today.Before(tomorrow),  // true
		today.Equal(todayAgain), // false (but very close)
	)
}

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
	// TODO...
}
