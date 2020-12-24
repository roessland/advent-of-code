package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBla(t *testing.T) {
	require.Equal(t, Dir{0,0,0}, Sum(Split("nwwswee")).Canonicalize())

	require.Equal(t, Dir{1,0, 0}, Split("esew")[0])
	require.Equal(t, Dir{0,0, 1}, Split("esew")[1])
	require.Equal(t, Dir{-1,0, 0}, Split("esew")[2])
	require.Equal(t, Dir{0,0,1}, Sum(Split("esew")))
	require.Equal(t, Dir{1,-1,0}, Sum(Split("esew")).Canonicalize())


}