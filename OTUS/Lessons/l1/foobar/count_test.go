package foobar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	s := "qwerasdfe"

	require.Equal(t, Count(s, 'e'), 2, "counting 'e' in "+s)
	require.Equal(t, Count(s, 'x'), 0, "counting 's' in "+s)
	require.Equal(t, Count(s, 'f'), 1, "counting 'f' in "+s)
}
