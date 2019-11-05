package uhelpers

import (
	"fmt"
	"strings"
)

func StringIndexOf(slice []string, item string) (int, error) {
	for index, iteratedItem := range slice {
		if strings.Compare(iteratedItem, item) == 0 {
			return index, nil
		}
	}

	return -1, fmt.Errorf("item not found in slice")
}
