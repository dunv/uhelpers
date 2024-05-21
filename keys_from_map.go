package uhelpers

func KeysFromMap(m map[string]any) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
	}
	return keys
}
