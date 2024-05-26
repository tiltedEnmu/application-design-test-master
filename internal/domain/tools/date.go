package tools

import "time"

func TimeToDate(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func NewDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DaysBetween(from, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := TimeToDate(from); !d.After(TimeToDate(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}
