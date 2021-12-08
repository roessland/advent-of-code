package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDist2(t *testing.T) {
	require.Equal(t, 0, Dist2([]int{1}, 1))
	require.Equal(t, 1, Dist2([]int{1}, 2))
	require.Equal(t, 3, Dist2([]int{1}, 3))
	require.Equal(t, 6, Dist2([]int{1}, 4))



}
