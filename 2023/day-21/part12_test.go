package main

import (
	"fmt"
	"testing"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountSeriesEven(t *testing.T) {
	t.Run("simple 1a", func(t *testing.T) {
		// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 (5 even)
		even := CountSeriesEven(1, 1, 10)
		assert.Equal(t, 5, even) // 2, 4, 6, 8, 10
	})

	t.Run("simple 1b", func(t *testing.T) {
		// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 (5 even)
		even := CountSeriesEven(1, 1, 11)
		assert.Equal(t, 5, even) // 2, 4, 6, 8, 10
	})

	t.Run("simple 2a", func(t *testing.T) {
		// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 (5 even)
		even := CountSeriesEven(0, 1, 9)
		assert.Equal(t, 5, even) // 0, 2, 4, 6, 8
	})

	t.Run("simple 2b", func(t *testing.T) {
		// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 (6 even)
		even := CountSeriesEven(0, 1, 10)
		assert.Equal(t, 6, even) // 0, 2, 4, 6, 8, 10
	})

	t.Run("simple 3a", func(t *testing.T) {
		// 0, 2, 4, 6, 8 (5 even)
		// 9-0 = 9
		// 9 / 2 = 4
		even := CountSeriesEven(0, 2, 9)
		assert.Equal(t, 5, even) // 0, 2, 4, 6, 8
	})

	t.Run("simple 3b", func(t *testing.T) {
		// 0, 2, 4, 6, 8, 10 (6 even)
		// 10-0 = 10
		// 10 / 2 = 5
		// But +1 since 2 divides 10
		even := CountSeriesEven(0, 2, 10)
		assert.Equal(t, 6, even) // 0, 2, 4, 6, 8, 10
	})

	t.Run("simple 4a", func(t *testing.T) {
		// 3, 7, 11, 15, 19 (0 even)
		assert.Equal(t, 0, CountSeriesEven(3, 4, 19))

		// 4, 8, 12, 16, 20 (5 even)
		assert.Equal(t, 4, CountSeriesEven(4, 4, 17)) // 4, 8, 12, 16
		assert.Equal(t, 4, CountSeriesEven(4, 4, 18)) // 4, 8, 12, 16
		assert.Equal(t, 4, CountSeriesEven(4, 4, 19)) // 4, 8, 12, 16
		assert.Equal(t, 5, CountSeriesEven(4, 4, 20)) // 4, 8, 12, 16, 20
		assert.Equal(t, 5, CountSeriesEven(4, 4, 21)) // 4, 8, 12, 16, 20
		assert.Equal(t, 5, CountSeriesEven(4, 4, 22)) // 4, 8, 12, 16, 20
		assert.Equal(t, 5, CountSeriesEven(4, 4, 23)) // 4, 8, 12, 16, 20
		assert.Equal(t, 6, CountSeriesEven(4, 4, 24)) // 4, 8, 12, 16, 20, 24
	})
}

func TestCountSeriesEvenOdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 (5 even, 5 odd)
		// => 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 (5 even, 5 odd)
		even := CountSeriesEvenOdd(1, 1, 10, false)
		odd := CountSeriesEvenOdd(1, 1, 10, true)
		assert.Equal(t, 5, even) // 2, 4, 6, 8, 10

		// 10 - 1 = 9
		// (9 + 1) / 2 = 5
		assert.Equal(t, 5, odd) // 1, 3, 5, 7, 9
	})

	t.Run("max below start", func(t *testing.T) {
		even := CountSeriesEvenOdd(313337, 1, 1337, false)
		odd := CountSeriesEvenOdd(313337, 1, 1337, true)
		assert.Equal(t, 0, even)
		assert.Equal(t, 0, odd)
	})

	t.Run("even interval", func(t *testing.T) {
		// 1, 3, 5, 7, 9
		even := CountSeriesEvenOdd(1, 2, 10, false)
		odd := CountSeriesEvenOdd(1, 2, 10, true)
		assert.Equal(t, 0, even)

		// 10 - 1 = 9
		// (9+1) / 2 = 5
		assert.Equal(t, 5, odd)
	})

	t.Run("simple 4a", func(t *testing.T) {
		// 3, 7, 11, 15, 19 (0 even)
		assert.Equal(t, 0, CountSeriesEven(3, 4, 19))

		// 3, 7, 11, 15, 19 (5 odd)
		assert.Equal(t, 4, CountSeriesEvenOdd(3, 4, 16, true)) // 3, 7, 11, 15
		assert.Equal(t, 4, CountSeriesEvenOdd(3, 4, 17, true)) // 3, 7, 11, 15
		assert.Equal(t, 4, CountSeriesEvenOdd(3, 4, 18, true)) // 3, 7, 11, 15
		assert.Equal(t, 5, CountSeriesEvenOdd(3, 4, 19, true)) // 3, 7, 11, 15, 19
		assert.Equal(t, 5, CountSeriesEvenOdd(3, 4, 20, true)) // 3, 7, 11, 15, 19
		assert.Equal(t, 5, CountSeriesEvenOdd(3, 4, 21, true)) // 3, 7, 11, 15, 19
		assert.Equal(t, 5, CountSeriesEvenOdd(3, 4, 22, true)) // 3, 7, 11, 15, 19
		assert.Equal(t, 6, CountSeriesEvenOdd(3, 4, 23, true)) // 3, 7, 11, 15, 19, 23
	})
}

func ReadTestInput() Map {
	return aocutil.ReadLines("input-ex1.txt")
}

func TestPickTileInQuadrantAtRadius(t *testing.T) {
	e0 := PickTileInQuadrantAtRadius(E, 0)
	assert.Equal(t, Tile{0, 0}, *e0)

	o0 := PickTileInQuadrantAtRadius(Origin, 0)
	assert.Equal(t, Tile{0, 0}, *o0)

	e1 := PickTileInQuadrantAtRadius(E, 1)
	assert.Equal(t, Tile{0, 1}, *e1)

	ne1 := PickTileInQuadrantAtRadius(NE, 1)
	assert.Nil(t, ne1)

	ne2 := PickTileInQuadrantAtRadius(NE, 2)
	assert.Equal(t, Tile{-1, 1}, *ne2)

	sw4 := PickTileInQuadrantAtRadius(SW, 4)
	assert.Equal(t, Tile{1, -3}, *sw4)

	se6 := PickTileInQuadrantAtRadius(SE, 6)
	assert.Equal(t, Tile{1, 5}, *se6)

	nw8 := PickTileInQuadrantAtRadius(NW, 8)
	assert.Equal(t, Tile{-7, -1}, *nw8)
}

// func NumReachablePosQuadrantInExactly(m *Map, cachedBFS2 map[Pos]int, S, p0 Pos, exactSteps int, quadrant Quadrant) int {
func TestNumReachablePosInQuadrantInExactly(t *testing.T) {
	m := ReadTestInput()
	// m := ReadInput()
	start := m.FindStartPos()
	cachedBFS2 := BFS2(&m, start, 15*m.SizeX())

	// From origin to origin
	{
		// East
		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 0, E)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 1, E)
			assert.Equal(t, 0, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 2, E)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 2, E)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 15, E)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 30, E)
			assert.Equal(t, 2, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 54, E)
			assert.Equal(t, 3, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 56, E)
			assert.Equal(t, 3, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 76, E)
			assert.Equal(t, 4, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 76+123456*22, E)
			assert.Equal(t, 4+123456, num)
		}
	}

	{
		{
			// South west
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 0, SW)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 1, SW)
			assert.Equal(t, 0, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 20, SW)
			assert.Equal(t, 1, num)
		}

		{
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 22, SW)
			require.Equal(t, 2, num)
		}

		{
			// x,,,,,,
			// ,x,x,,,
			// ,,x,,,,
			// ,x,,,,,
			num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 22+22, SW)
			require.Equal(t, 5, num)
		}
	}

	// Multiple tiles in quadrant should be summed
	{
		// x,,,,,,
		// ,x,,,,,
		// ,,,,,,,
		// ,,,,,,,
		p := Pos{m.SizeY() - 1, m.SizeX() - 1}
		num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, p, 36, SE)
		require.Equal(t, 2, num)
	}

	// Multiple tiles in quadrant should be summed
	{
		// x,,,,,,
		// ,x,x,,,
		// ,,x,,,,
		// ,x,,,,,
		p := Pos{m.SizeY() - 1, m.SizeX() - 1}
		num := NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, p, 36+22, SE)
		require.Equal(t, 5, num)
	}
}

func TestNumReachablePosInExactly(t *testing.T) {
	m := ReadTestInput()
	start := m.FindStartPos()
	cachedBFS2 := BFS2(&m, start, 15*m.SizeX())

	{
		num := NumReachablePosInExactly(&m, cachedBFS2, start, start, 0)
		assert.Equal(t, 1, num)
	}
	{
		num := NumReachablePosInExactly(&m, cachedBFS2, start, start, 1)
		assert.Equal(t, 0, num)
	}
	{
		num := NumReachablePosInExactly(&m, cachedBFS2, start, start, 2)
		assert.Equal(t, 1, num)
	}

	{

		// The start in nearby tiles is reachable at these steps (even):
		// 0, 30, 22, 32, 26, 30, 22, 32, 26
		assert.Equal(t, 1, NumReachablePosInExactly(&m, cachedBFS2, start, start, 0))
		assert.Equal(t, 3, NumReachablePosInExactly(&m, cachedBFS2, start, start, 22))
		assert.Equal(t, 5, NumReachablePosInExactly(&m, cachedBFS2, start, start, 26))
		assert.Equal(t, 7, NumReachablePosInExactly(&m, cachedBFS2, start, start, 30))
		assert.Equal(t, 9, NumReachablePosInExactly(&m, cachedBFS2, start, start, 32))

		// The start in nearby tiles is reachable at these steps (odd):
		// 15, 33, 21, 37, 15, 33, 21, 37
		assert.Equal(t, 0, NumReachablePosInExactly(&m, cachedBFS2, start, start, 1))
		assert.Equal(t, 2, NumReachablePosInExactly(&m, cachedBFS2, start, start, 15))
		assert.Equal(t, 4, NumReachablePosInExactly(&m, cachedBFS2, start, start, 21))
		assert.Equal(t, 8, NumReachablePosInExactly(&m, cachedBFS2, start, start, 33))
		assert.Equal(t, 12, NumReachablePosInExactly(&m, cachedBFS2, start, start, 37))
	}

	{

		// The origin in nearby tiles is reachable at these steps (even):
		// 10, 22, 22, 32, 32, 32, 22, 22, 16
		p0 := Pos{0, 0}
		assert.Equal(t, 0, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 0))
		assert.Equal(t, 1, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 10))
		assert.Equal(t, 2, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 16))
		assert.Equal(t, 6, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 22))
		assert.Equal(t, 9, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 32))

		// p0 in nearby tiles is reachable at these steps (odd):
		// Fst: 11, 33, 21, 43, 21, 33, 11, 27
		// Nxt: 33, 55, 43, 65, 43, 55, 33, 49
		//
		// ,,,,
		// ,,,,
		// ,,,,
		// ,,,,
		//
		assert.Equal(t, 0, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 1))
		assert.Equal(t, 2, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 11))
		assert.Equal(t, 4, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 21))
		assert.Equal(t, 6, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 27))
		assert.Equal(t, 12, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 33))
		assert.Equal(t, 16, NumReachablePosInExactly(&m, cachedBFS2, start, p0, 43))

		for _, quadrant := range AllQuadrants {
			NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, Pos{0, 0}, 22+50*22+1, quadrant)
		}
	}
}

func TestNumReachableExactly(t *testing.T) {
	m := ReadTestInput()
	start := m.FindStartPos()
	cachedBFS2 := BFS2(&m, start, 20*m.SizeX())
	assert.Equal(t, 1, NumReachableInExactly(&m, cachedBFS2, start, 0))
	assert.Equal(t, 2, NumReachableInExactly(&m, cachedBFS2, start, 1))
	assert.Equal(t, 4, NumReachableInExactly(&m, cachedBFS2, start, 2))

	for _, quadrant := range AllQuadrants {
		assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start, 4, quadrant))
	}
	assert.Equal(t, 6, NumReachableInExactly(&m, cachedBFS2, start, 3))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West(), 3, W))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.North(), 3, N))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.North().North().East(), 3, NE))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 3, NW))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West(), 3, SW))
	assert.Equal(t, 1, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West(), 3, SE))

	assert.Equal(t, 2, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 3+40, NW))
	assert.Equal(t, 2, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 25, NW))
	assert.Equal(t, 1+1+3, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47, NW))
	assert.Equal(t, 1+1+3+5, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47+22*1, NW))
	assert.Equal(t, 1+1+3+5+7, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47+22*2, NW))
	assert.Equal(t, 1+1+3+5+7+9, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47+22*3, NW))
	assert.Equal(t, 1+1+3+5+7+9+11, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47+22*4, NW))
	assert.Equal(t, 1+1+3+5+7+9+11+13, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), 47+22*5, NW))

	r := 7
	dr := 2
	a := 47 + 22*1
	da := 22
	expec := 1 + 1 + 3 + 5
	for i := 0; i < 5; i++ {
		expec += r
		r += dr
		a += da
		assert.Equal(t, expec, NumReachablePosQuadrantInExactly(&m, cachedBFS2, start, start.West().West().North(), a, NW), fmt.Sprintf("i=%d, a=%d", i, a))
	}

	assert.Equal(t, 9, NumReachableInExactly(&m, cachedBFS2, start, 4))
	assert.Equal(t, 13, NumReachableInExactly(&m, cachedBFS2, start, 5))
	assert.Equal(t, 16, NumReachableInExactly(&m, cachedBFS2, start, 6))
	assert.Equal(t, 50, NumReachableInExactly(&m, cachedBFS2, start, 10))
	assert.Equal(t, 1594, NumReachableInExactly(&m, cachedBFS2, start, 50))
	assert.Equal(t, 6536, NumReachableInExactly(&m, cachedBFS2, start, 100))
	assert.Equal(t, 167004, NumReachableInExactly(&m, cachedBFS2, start, 500))
	assert.Equal(t, 668697, NumReachableInExactly(&m, cachedBFS2, start, 1000))
	assert.Equal(t, 16733044, NumReachableInExactly(&m, cachedBFS2, start, 5000))
}

func TestNumReachableExactlyRealINput(t *testing.T) {
	m := ReadInput()
	start := m.FindStartPos()
	cachedBFS2 := BFS2(&m, start, 15*m.SizeX())
	assert.Equal(t, 592723929260582, NumReachableInExactly(&m, cachedBFS2, start, 26501365))
}

// func CountTilesInQuadrantAtRadius(quadrant Quadrant, radius int) int {
func TestCountTilesInQuadrantAtRadius(t *testing.T) {
	assert.Equal(t, 1, CountTilesInQuadrantAtRadius(SE, 0))
	assert.Equal(t, 1, CountTilesInQuadrantAtRadius(SE, 1))
	assert.Equal(t, 2, CountTilesInQuadrantAtRadius(SE, 3))
	assert.Equal(t, 3, CountTilesInQuadrantAtRadius(SE, 4))
	assert.Equal(t, 1, CountTilesInQuadrantAtRadius(E, 5))
}

func TestSumSeries(t *testing.T) {
	// ,,,,,
	// ,,x,x
	// .x,x,
	// ,,x,,
	// ,x,,,
	// Steps to first diagonal is a0
	// Steps to next diagonal is a0 + da
	// Length of first diagonal is r.
	// Length of next diagonal is r + dr
	// Count number of x's where the number of steps is below max.

	{
		a0 := 10
		da := 21
		r0 := 2
		dr := 2
		max := a0 + da
		assert.Equal(t, 6, SumSeries(a0, da, r0, dr, max))
	}

	{
		// ,,,,,,,
		// ,,x,x,x
		// .x,x,x,
		// ,,x,x,,
		// ,x,x,,,
		// ,,x,,,,
		// ,x,,,,,
		//
		a0 := 10
		da := 21
		r0 := 2
		dr := 2
		max := a0 + da + da
		// max - a0 = 2da = 42
		//  (max-a0)/da = 2
		assert.Equal(t, 12, SumSeries(a0, da, r0, dr, max))
	}

	{
		// x
		//
		a0 := 5
		da := 22
		r0 := 1
		dr := 0
		max := 5
		// max - a0 = 2da = 42
		//  (max-a0)/da = 2
		assert.Equal(t, 1, SumSeries(a0, da, r0, dr, max))
	}

	{
		// ,,,,x,x,x,x
		// s = 5, 27, 49, 71, ...
		// r = 1, 1, 1, 1, ...
		a0 := 5
		da := 22
		r0 := 1
		dr := 0
		max := 27
		assert.Equal(t, 2, SumSeries(a0, da, r0, dr, max))
		assert.Equal(t, 2, SumSeries(a0, da, r0, dr, max+1))
		assert.Equal(t, 1, SumSeries(a0, da, r0, dr, max-1))
	}

	{
		// ,,,,x,x,x,x
		// s = 5, 27, 49, 71, ...
		// r = 1, 1, 1, 1, ...
		a0 := 142
		da := 22
		r0 := 1
		dr := 0
		max := 0
		// max - a0 = -142
		assert.Equal(t, 0, SumSeries(a0, da, r0, dr, max))
	}

	{
		a0 := 135
		da := 22
		r0 := 9
		dr := 2
		max := 157
		// 135, 157 are less than equal to 157.
		// So we expect 9 + 11 = 20
		assert.Equal(t, 20, SumSeries(a0, da, r0, dr, max))
	}

	{
		a0 := 135
		da := 22
		r0 := 9
		dr := 2
		max := 135
		// 135 is less than equal to 135
		// So we expect 9
		assert.Equal(t, 9, SumSeries(a0, da, r0, dr, max))
	}
}
