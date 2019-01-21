package helpers

import "time"

// PtrToString <-
func PtrToString(s string) *string {
	return &s
}

// PtrToTime <-
func PtrToTime(s time.Time) *time.Time {
	return &s
}
