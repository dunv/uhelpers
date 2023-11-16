package uhelpers_test

import (
	"testing"
	"time"

	"github.com/dunv/uhelpers"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	stringManual := "test"
	stringPtr := uhelpers.Ptr(stringManual)
	require.Equal(t, stringPtr, &stringManual)

	timeManual := time.Now()
	timePtr := uhelpers.Ptr(timeManual)
	require.Equal(t, timePtr, &timeManual)

	boolManual := time.Now()
	boolPtr := uhelpers.Ptr(boolManual)
	require.Equal(t, boolPtr, &boolManual)

}
