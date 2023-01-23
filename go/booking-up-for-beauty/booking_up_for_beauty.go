package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:04"
	t, _ := time.Parse(layout, date)

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	result := false

	layout := "Monday, January 2, 2006 15:04:04"
	t, _ := time.Parse(layout, date)

	hour := t.Hour()
	if hour >= 12 && hour < 18 {
		result = true
	}

	return result
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:00"
	t, _ := time.Parse(layout, date)

	formattedDate := t.Format("Monday, January 2, 2006, at 15:04")

	result := fmt.Sprintf("You have an appointment on %s.", formattedDate)
	return result
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	const (
		openingDay   = 15
		openingMonth = 9
		openingYear  = 2021
	)
	year := time.Now().Year()
	t := time.Date(year, openingMonth, openingDay, 0, 0, 0, 0, time.UTC)
	return t
}
