package uhelpers

import (
	"fmt"
	"time"
)

// Get first hour of a day
func FirstHourOfDay(reference time.Time) time.Time {
	return time.Date(reference.Year(), reference.Month(), reference.Day(), 0, 0, 0, 0, reference.Location())
}

// Figure out which day is the first of an ISOWeek
func FirstDayOfISOWeek(reference time.Time) time.Time {
	year, week := reference.ISOWeek()
	timezone := reference.Location()

	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		// isoYear, isoWeek = date.ISOWeek()
	}

	return date
}

// Get the first day of the month
func FirstDayOfMonth(reference time.Time) time.Time {
	return time.Date(reference.Year(), reference.Month(), 0, 0, 0, 0, 0, reference.Location())
}

// Get the first day of a year
func FirstDayOfYear(reference time.Time) time.Time {
	return time.Date(reference.Year(), 0, 0, 0, 0, 0, 0, reference.Location())
}

// Generate a time.Time from year and mont
func TimeFromMonthAndYear(year int, month int, timeZone string) (*time.Time, error) {
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, fmt.Errorf("Could not load timezone %s. (%v)", timeZone, err)
	}

	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	return &date, nil
}

// Check if two time.Time are on the same day
func IsSameDayWithoutTimezoneSupport(t time.Time, other time.Time) bool {
	return t.Year() == other.Year() && t.YearDay() == other.YearDay()
}
