package uhelpers

func MapStringArray(in []string, f func(in string) (string, error)) ([]string, error) {
	result := []string{}
	for _, elem := range in {
		formatted, err := f(elem)
		if err != nil {
			return nil, err
		}
		result = append(result, formatted)
	}
	return result, nil
}
