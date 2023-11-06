package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	fmtDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Now()
	}
	return fmtDate.UTC()
}

// HasPassed returns whether a date has passed
// Must use `January 2, 2006 15:04:05` as the date mark, otherwise the format fails
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"

	fmtDate, err := time.Parse(layout, date)
	if err != nil {
		return false
	}

	isPassed := time.Now().After(fmtDate)
	return isPassed
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"

	fmtDate, err := time.Parse(layout, date)
	if err != nil {
	}
	if fmtDate.Hour() < 18 && fmtDate.Hour() >= 12 {
		return true
	}
	return false

}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"

	fmtDate, err := time.Parse(layout, date)
	if err != nil {
	}
	h, m, _ := fmtDate.Clock()
	year, _, day := fmtDate.Date()
	weekday := fmtDate.Weekday().String()
	month := fmtDate.Month().String()

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", weekday, month, day, year, h, m)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	const layout = "2006/01/02 15:04:05"
	const date = "2023/09/15 00:00:00"

	fmtDate, err := time.Parse(layout, date)
	if err != nil {
	}
	return fmtDate.UTC()
}
