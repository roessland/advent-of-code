package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"
)

func TestFindVertex(t *testing.T) {
	{
		// >>>
		// ..v
		a := Line{Pos{0, 0}, Pos{0, 2}}
		b := Line{Pos{0, 2}, Pos{1, 2}}
		require.Equal(t, Pos{1, 2}, findVertex(a, b))
	}

	{
		//   ^
		// >>>
		a := Line{Pos{1, 0}, Pos{1, 2}}
		b := Line{Pos{1, 2}, Pos{0, 2}}
		require.Equal(t, Pos{2, 3}, findVertex(a, b))
	}
}

func TestFindVertices(t *testing.T) {
	{
		// >>>>
		// ^..v
		// ^..v
		// <<<<
		lines := []Line{
			{Pos{0, 0}, Pos{0, 3}},
			{Pos{0, 3}, Pos{3, 3}},
			{Pos{3, 3}, Pos{3, 0}},
			{Pos{3, 0}, Pos{0, 0}},
		}
		vertices := findVertices(lines)
		require.EqualValues(t, []Pos{
			{1, 3},
			{3, 3},
			{3, 1},
			{1, 1},
		}, vertices)
	}
	{
		// <<<<
		// v..^
		// v..^
		// >>>>
		lines := []Line{
			{Pos{0, 3}, Pos{0, 0}},
			{Pos{0, 0}, Pos{3, 0}},
			{Pos{3, 0}, Pos{3, 3}},
			{Pos{3, 3}, Pos{0, 3}},
		}
		vertices := findVertices(lines)
		require.EqualValues(t, []Pos{
			{0, 0},
			{4, 0},
			{4, 4},
			{0, 4},
		}, vertices)
	}
}

func TestMeasureTrenchSize(t *testing.T) {
	{
		// ...
		// >>>
		// ^.v
		// <<<
		lines := []Line{
			{Pos{1, 0}, Pos{1, 2}}, // -3
			{Pos{1, 2}, Pos{3, 2}}, // 0?
			{Pos{3, 2}, Pos{3, 0}}, // +12
			{Pos{3, 0}, Pos{1, 0}}, // 0?
		}
		slices.Reverse(lines)
		for i := range lines {
			lines[i] = lines[i].Reversed()
		}
		require.Equal(t, 9, measureTrenchSize(lines))
	}

	{
		// ....
		// >>>>
		// ^..v
		// <<v<
		//  ^v
		lines := []Line{
			{Pos{1, 0}, Pos{1, 3}}, // -4
			{Pos{1, 3}, Pos{4, 3}}, // 0?
			{Pos{3, 3}, Pos{3, 2}}, // 8
			{Pos{3, 2}, Pos{4, 2}}, // 0?
			{Pos{4, 2}, Pos{4, 1}},
			{Pos{4, 1}, Pos{3, 1}}, // 0?
			{Pos{3, 1}, Pos{3, 0}}, // 8
			{Pos{3, 0}, Pos{1, 0}}, // 0?
		}
		slices.Reverse(lines)
		for i := range lines {
			lines[i] = lines[i].Reversed()
		}
		require.Equal(t, 14, measureTrenchSize(lines))
	}
}
