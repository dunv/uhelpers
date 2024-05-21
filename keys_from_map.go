package uhelpers

func KeysFromMap[T any, K comparable](m map[K]T) []K {
	keys := make([]K, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}
