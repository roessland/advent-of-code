package day04_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2025/day-04/day04"
	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	inputEx := day04.ReadInput("input-ex1.txt")
	require.Equal(t, 13, day04.Part1(inputEx))
	require.Equal(t, 43, day04.Part2(inputEx))

	input := day04.ReadInput("input.txt")
	require.Equal(t, 1523, day04.Part1(input))
	require.Equal(t, 9290, day04.Part2(input))
}
