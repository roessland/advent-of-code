package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanGoTo(t *testing.T) {
	require.True(t, HasKeyToDoor('g', make(map[rune]struct{})))
	require.False(t, HasKeyToDoor('G', make(map[rune]struct{})))
	require.True(t, HasKeyToDoor('G', map[rune]struct{}{'g': {}}))
}

func TestWithKey(t *testing.T) {
	before := map[rune]struct{}{'g': {}, 'b': {}}
	after := WithKey(before, 'a')
	require.Len(t, before, 2)
	require.Len(t, after, 3)
	require.Contains(t, after, 'a')
	require.Contains(t, after, 'g')
	require.Contains(t, after, 'b')


	after2 := WithKey(after, 'A')
	require.Len(t, after2, 3)
	require.Contains(t, after2, 'a')
	require.Contains(t, after2, 'g')
	require.Contains(t, after2, 'b')


}
