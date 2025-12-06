package day03_test

import (
	"testing"

	"github.com/roessland/advent-of-code/2025/day-03/day03"
	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	input := day03.ReadInput("input-ex1.txt")
	require.Equal(t, input[0][0], 9)
	require.Len(t, input, 4)
}

func TestLargestJoltPossible2(t *testing.T) {
	f := day03.LargestJoltPossible2
	require.Equal(t, 12, f([]int{1, 2}))
	require.Equal(t, 21, f([]int{2, 1}))
	require.Equal(t, 91, f([]int{8, 9, 1}))
}

func TestPart1(t *testing.T) {
	input := day03.ReadInput("input.txt")
	require.Equal(t, 17435, day03.Part1(input))
}

func TestLargestJoltPossibleN(t *testing.T) {
	f := day03.LargestJoltPossibleN
	require.Equal(t, 12, f([]int{1, 2}, 2))
	require.Equal(t, 21, f([]int{2, 1}, 2))
	require.Equal(t, 91, f([]int{8, 9, 1}, 2))
	require.Equal(t, 891, f([]int{8, 9, 1}, 3))
	require.Equal(t, 891, f([]int{1, 8, 9, 1}, 3))
	require.Equal(t, 329, f([]int{3, 1, 2, 9}, 3))
	require.Equal(t, 3129, f([]int{3, 1, 2, 9}, 4))
	require.Equal(t, 434234234278, f([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 12))
}

func TestPart2(t *testing.T) {
	input_ := day03.ReadInput("input-ex1.txt")
	require.Equal(t, 3121910778619, day03.Part2(input_))

	input := day03.ReadInput("input.txt")
	require.Equal(t, 172886048065379, day03.Part2(input))
}
