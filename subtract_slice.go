package uhelpers

// Returns a slice which contains all elements of a that are not in b
func SubtractSlice[T comparable](a []T, b []T) []T {
	result := []T{}
	for i := range a {
		elemA := a[i]
		found := false
		for j := range b {
			elemB := b[j]
			if elemA == elemB {
				found = true
				break
			}
		}
		if !found {
			result = append(result, elemA)
		}
	}
	return result
}
