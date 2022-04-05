package uhelpers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	stringManual := "test"
	stringPtr := Ptr(stringManual)
	require.Equal(t, stringPtr, &stringManual)

	timeManual := time.Now()
	timePtr := Ptr(timeManual)
	require.Equal(t, timePtr, &timeManual)

	boolManual := time.Now()
	boolPtr := Ptr(boolManual)
	require.Equal(t, boolPtr, &boolManual)

}
