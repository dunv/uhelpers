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

func PtrToInt(s int) *int {
	return &s
}

func PtrToInt32(s int32) *int32 {
	return &s
}

func PtrToInt64(s int64) *int64 {
	return &s
}

func PtrToFloat32(s float32) *float32 {
	return &s
}

func PtrToFloat64(s float64) *float64 {
	return &s
}

func PtrToDuration(s time.Duration) *time.Duration {
	return &s
}
