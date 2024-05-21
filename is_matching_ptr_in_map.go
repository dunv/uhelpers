package uhelpers

func IsMatchingStringPointerInMap(ptr1 *string, rawMap interface{}, key string, bothNilReturnTrue ...bool) bool {
	if m, ok := rawMap.(map[string]interface{}); ok {
		if rawVal, ok := m[key]; ok {
			if val, ok := rawVal.(*string); ok {
				return CompareStringPointers(ptr1, val, bothNilReturnTrue...)
			}
		}
	}
	return false
}

func IsMatchingBoolPointerInMap(ptr1 *bool, rawMap interface{}, key string, bothNilReturnTrue ...bool) bool {
	if m, ok := rawMap.(map[string]interface{}); ok {
		if rawVal, ok := m[key]; ok {
			if val, ok := rawVal.(*bool); ok {
				return CompareBoolPointers(ptr1, val, bothNilReturnTrue...)
			}
		}
	}
	return false
}
