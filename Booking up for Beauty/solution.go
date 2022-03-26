package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}
// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
		t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	m := t.Format("Monday, January 2, 2006, at 15:04.")
	s := fmt.Sprint("You have an appointment on ", m)
	return s
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
