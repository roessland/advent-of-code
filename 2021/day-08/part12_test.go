package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSet(t *testing.T) {
	require.EqualValues(t, 0b0, Signals(0))
	require.EqualValues(t, 0b1, SetOf("a"))
	require.EqualValues(t, 0b10, SetOf("b"))
	require.EqualValues(t, 0b11, SetOf("ba"))
	require.EqualValues(t, 0b101, SetOf("c")|SetOf("a"))
	require.EqualValues(t, 5, SetOf("acdeg").Len())
	require.EqualValues(t, 2, (SetOf("acdeg") & SetOf("acf")).Len())
	require.EqualValues(t, 6, (SetOf("acdeg") | SetOf("acf")).Len())
}
