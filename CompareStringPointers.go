package uhelpers

func CompareStringPointers(ptr1 *string, ptr2 *string, bothNilReturnTrue ...bool) bool {
	if ptr1 == nil && ptr2 == nil {
		if len(bothNilReturnTrue) > 0 {
			return bothNilReturnTrue[0]
		}
		return false
	}

	if ptr1 != nil && ptr2 == nil {
		return false
	}

	if ptr1 == nil && ptr2 != nil {
		return false
	}

	if *ptr1 == *ptr2 {
		return true
	}

	return false
}
