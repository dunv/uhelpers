package uhelpers_test

import (
	"testing"

	"github.com/dunv/uhelpers"
	"github.com/stretchr/testify/require"
)

func TestKeysFromMapString(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	keys := uhelpers.KeysFromMap(m)
	require.Contains(t, keys, "one")
	require.Contains(t, keys, "two")
	require.Contains(t, keys, "three")
}

func TestKeysFromMapCustomType(t *testing.T) {
	type testType string
	one := testType("one")
	two := testType("two")
	three := testType("three")
	m := map[testType]int{
		one:   1,
		two:   2,
		three: 3,
	}
	keys := uhelpers.KeysFromMap(m)
	require.Contains(t, keys, one)
	require.Contains(t, keys, two)
	require.Contains(t, keys, three)
}
