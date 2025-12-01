package day01_test

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/roessland/advent-of-code/2025/day-01/day01"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	input := day01.ReadInput("input-ex1.txt")
	require.Equal(t, input[0], -68)
	require.Equal(t, len(input), 10)
	require.Equal(t, input[8], 14)
}

func TestCumSum(t *testing.T) {
	cumSum := day01.CumSumMod([]int{1, 2, 3, 99}, 100)
	require.EqualValues(t, []int{1, 3, 6, 5}, cumSum)
}

func TestPart1(t *testing.T) {
	inputEx1 := day01.ReadInput("input-ex1.txt")
	require.Equal(t, 3, day01.Part1(inputEx1))

	input := day01.ReadInput("input.txt")
	require.Equal(t, 982, day01.Part1(input))
}

func TestModInterval(t *testing.T) {
	check := func(in1, in2, e1, e2 int) {
		a1, a2 := day01.ModInterval(in1, in2)
		assert.Equal(t, e1, a1)
		assert.Equal(t, e2, a2)
	}
	check(0, 50, 0, 50)
	check(-50, 0, -50, 0)
	check(-150, -50, -50, 50)
}

func TestVisits(t *testing.T) {
	type TC struct {
		Name     string
		Start    int
		Rot      int
		Expected day01.Expl
	}
	tcs := []TC{
		{
			"Full rotation R", 50, 100,
			day01.Expl{FullRotations: 1},
		},
		{
			"2 full rotations R", 50, 200,
			day01.Expl{FullRotations: 2},
		},
		{
			"Full rotation L", 50, -100,
			day01.Expl{FullRotations: 1},
		},
		{
			"2 full rotations L", 50, -200,
			day01.Expl{FullRotations: 2},
		},
		{
			"2 full rotations L", 50, -200,
			day01.Expl{FullRotations: 2},
		},
		{
			"almost R", 1, 98,
			day01.Expl{},
		},
		{
			"almost L", 99, -98,
			day01.Expl{},
		},
		{
			"exact R", 99, 1,
			day01.Expl{EndedUp: 1},
		},
		{
			"exact L", 1, -1,
			day01.Expl{EndedUp: 1},
		},
		{
			"pass R", 99, 2,
			day01.Expl{Passthroughs: 1},
		},
		{
			"pass L", 1, -2,
			day01.Expl{Passthroughs: 1},
		},
		{
			"full + pass L", 1, -2 - 100,
			day01.Expl{Passthroughs: 1, FullRotations: 1},
		},
		{
			"2 full L + exact L", 99, -99 - 200,
			day01.Expl{EndedUp: 1, FullRotations: 2},
		},
		{
			"exact R", 50, 50,
			day01.Expl{EndedUp: 1},
		},
		// The dial is rotated -1 to point at 99, during this rotation it points at zero 1 times
		{
			"from 0 L", 0, -1,
			day01.Expl{},
		},
		// The dial at a=14 is rotated -82 to point at b=32, during this rotation it points at zero 1 times
		{
			"from + to -", 14, -82,
			day01.Expl{Passthroughs: 1},
		},
		// 2025/12/01 21:11:54 INFO moved from=49 rot=90 to=39 visits=0
		{
			"from 2nd quadrant + to 2nd quadrant", 49, 90,
			day01.Expl{Passthroughs: 1},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			actual := day01.VisitsVerbose(tc.Start, tc.Rot)
			require.EqualValues(t, tc.Expected, actual, "expected != fast")

			actualSlow := day01.VisitsSlow(tc.Start, tc.Rot)
			require.EqualValues(t, tc.Expected.Total(), actualSlow, "slow not as expected")

			msg := fmt.Sprintf("slow != fast (slow=%d, fast=%#v)", actualSlow, actual)
			require.EqualValues(t, actualSlow, actual.Total(), msg)
		})
	}
}

func TestPart2(t *testing.T) {
	inputEx1 := day01.ReadInput("input-ex1.txt")
	require.Equal(t, 6, day01.Part2(inputEx1))

	input := day01.ReadInput("input.txt")
	t0 := time.Now()
	require.Equal(t, 6106, day01.Part2(input))

	fmt.Println("timing", slog.Duration("duration", time.Since(t0)))
}
