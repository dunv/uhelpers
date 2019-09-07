package uhelpers

import "time"

func PtrToString(s string) *string {
	return &s
}

func PtrToTime(s time.Time) *time.Time {
	return &s
}

func PtrToBool(s bool) *bool {
	return &s
}
