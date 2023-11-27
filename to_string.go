package uhelpers

import "fmt"

func ToString[T any](in *T) string {
	if in == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", *in)
}
