package booking

import "time"

// This is one of the most bizarre language-design choices
// I've ever seen in my life.

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	parsed, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err) // We haven't been taught this in the course yet?
	}
	return parsed
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	parsed, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	parsed, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsed.Hour() >= 12 && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parsed, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return "You have an appointment on " + parsed.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// In better coding, this would be based on a static original_date.
	// But Go's date parsing has made me lose all faith in love and beauty.
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
