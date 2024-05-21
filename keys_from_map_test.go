package uhelpers_test

import (
	"testing"

	"github.com/dunv/uhelpers"
	"github.com/stretchr/testify/require"
)

func TestKeysFromMap(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	keys := uhelpers.KeysFromMap(m)
	require.Equal(t, []string{"one", "two", "three"}, keys)
}
