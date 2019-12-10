package uhelpers

func SubtractStringArray(a []string, b []string) []string {
	result := []string{}
	for _, elem := range a {
		if _, err := StringIndexOf(b, elem); err != nil {
			result = append(result, elem)
		}
	}
	return result
}
