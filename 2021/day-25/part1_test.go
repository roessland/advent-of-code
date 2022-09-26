package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRem(t *testing.T) {
	require.Equal(t, 138, rem(-1, 139))
	require.Equal(t, 138, rem(-140, 139))
	require.Equal(t, 1, rem(140, 139))
}

func TestAlice(t *testing.T) {
	s := []byte{1, 2, 3}
	var a [4]byte

	require.Panics(t, func() {
		a = *(*[4]byte)(s)
		_ = a
	})
}
