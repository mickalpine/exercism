package booking

import (
	"log"
	"time"
)

const (
	scheduleLayout    string = "1/2/2006 15:04:05"
	passedLayout      string = "January 2, 2006 15:04:05"
	afternoonLayout   string = "Monday, January 2, 2006 15:04:05"
	descriptionLayout string = "You have an appointment on Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse(scheduleLayout, date)
	if err != nil {
		log.Fatalf("Unable to parse the time, ensure format is: %s", scheduleLayout)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse(passedLayout, date)
	if err != nil {
		panic(err)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse(afternoonLayout, date)
	if err != nil {
		panic(err)
	}

	hour := t.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return Schedule(date).Format(descriptionLayout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
