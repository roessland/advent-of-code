package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNextStableWaterPropagatesHitLeftWall(t *testing.T) {
	//#....|#.....
	//#~||||#.....
	//#######.....
	prev := Tile{Type: FallingWater}
	up := Tile{Type: Sand}
	down := Tile{Type: Clay}
	left := Tile{Type: StableWater}
	right := Tile{Type: FallingWater}
	next := NextTile(prev, up, down, left, right, Tile{}, Tile{})
	require.True(t, next.ReachedWallLeft)
}

func TestNextFallingWaterBecomesStable(t *testing.T) {
	//#....|#.....
	//#~~~~|#.....
	//#######.....
	prev := Tile{Type: FallingWater, VelX: -1}
	up := Tile{Type: FallingWater}
	down := Tile{Type: Clay}
	left := Tile{Type: StableWater}
	right := Tile{Type: Clay}
	next := NextTile(prev, up, down, left, right, Tile{}, Tile{})
	require.Equal(t, StableWater, next.Type)
}

func TestNextFallingWaterChoosesDirection(t *testing.T) {
	//.....|.....#
	//#..#.|.....#
	//#..#~~#.....
	//#..#~~#.....
	prev := Tile{Type: FallingWater}
	up := Tile{Type: FallingWater}
	down := Tile{Type: StableWater}
	left := Tile{Type: Sand}
	right := Tile{Type: Sand}
	next := NextTile(prev, up, down, left, right, Tile{}, Tile{})
	require.True(t, next.VelX != 0)
}

func TestNextFallingWaterChoosesDirection2(t *testing.T) {
	//.....|.....#
	//#..#@<.....#
	//#..#~~#.....
	//#..#~~#.....
	prev := Tile{Type: Sand}
	up := Tile{Type: Sand}
	down := Tile{Type: StableWater}
	left := Tile{Type: Clay}
	right := Tile{Type: FallingWater, VelX: -1}
	downright := Tile{Type: StableWater}
	next := NextTile(prev, up, down, left, right, Tile{}, downright)
	require.Equal(t, FallingWater, next.Type)
}

func TestNextFallingWaterChoosesDirection3(t *testing.T) {
	//.....|.....#
	//#..#.>@....#
	//#..#~~#.....
	//#..#~~#.....
	prev := Tile{Type: Sand}
	up := Tile{Type: Sand}
	down := Tile{Type: Clay}
	left := Tile{Type: FallingWater, VelX: 1}
	right := Tile{Type: Sand}
	downleft := Tile{Type: StableWater}
	next := NextTile(prev, up, down, left, right, downleft, Tile{})
	require.Equal(t, FallingWater, next.Type)
}

func TestNextFallingWaterStopsIfSandBelow(t *testing.T) {
	//.....|.....#
	//#..#.>.>@..#
	//#..#~~#.....
	//#..#~~#.....
	prev := Tile{Type: Sand}
	up := Tile{Type: Sand}
	down := Tile{Type: Sand}
	left := Tile{Type: FallingWater, VelX: 1}
	right := Tile{Type: Sand}
	downleft := Tile{Type: Sand}
	next := NextTile(prev, up, down, left, right, downleft, Tile{})
	require.Equal(t, Sand, next.Type)
}

func TestChangesDirection(t *testing.T) {
	//.....|.....#
	//#..#.>.>...#
	//#..#~~#.....

	prev := Tile{Type: FallingWater, VelX: 1}
	up := Tile{Type: FallingWater}
	down := Tile{Type: StableWater}
	left := Tile{Type: Sand}
	right := Tile{Type: Sand}
	downleft := Tile{Type: StableWater}
	downright := Tile{Type: Clay}

	hasleftcount := 0
	for i := 0; i < 10; i++ {
		if NextTile(prev, up, down, left, right, downleft, downright).VelX == -1 {
			hasleftcount++
		}
	}
	require.Greater(t, hasleftcount, 0)
}
