package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFall(t *testing.T) {
	bricks := []Brick{
		{Name: "A", Origin: Coord3{1, 0, 1}, Size: Coord3{1, 3, 1}},
		{Name: "B", Origin: Coord3{0, 0, 2}, Size: Coord3{3, 1, 1}},
		{Name: "C", Origin: Coord3{0, 2, 3}, Size: Coord3{3, 1, 1}},
		{Name: "D", Origin: Coord3{0, 0, 4}, Size: Coord3{1, 3, 1}},
		{Name: "E", Origin: Coord3{2, 0, 5}, Size: Coord3{1, 3, 1}},
		{Name: "F", Origin: Coord3{0, 1, 6}, Size: Coord3{3, 1, 1}},
		{Name: "G", Origin: Coord3{1, 1, 8}, Size: Coord3{1, 1, 2}},
		{Name: "💚", Origin: Coord3{0, 0, 0}, Size: Coord3{10, 10, 1}},
	}

	fallen := Fall(bricks)
	for _, b := range fallen {
		fmt.Printf("%+v\n", b)
	}

	byName := map[string]Brick{}
	for _, b := range fallen {
		byName[b.Name] = b
	}

	assert.Equal(t, 1, byName["A"].Origin.Z)
	assert.Equal(t, 2, byName["B"].Origin.Z)
	assert.Equal(t, 2, byName["C"].Origin.Z)
	assert.Equal(t, 3, byName["D"].Origin.Z)
	assert.Equal(t, 3, byName["E"].Origin.Z)
	assert.Equal(t, 4, byName["F"].Origin.Z)
	assert.Equal(t, 5, byName["G"].Origin.Z)
}

func TestContainsXY(t *testing.T) {
	br := Brick{Name: "B", Origin: Coord3{0, 0, 2}, Size: Coord3{3, 1, 1}}
	assert.False(t, br.ContainsXY(Coord2{-1, 0}))
	assert.True(t, br.ContainsXY(Coord2{0, 0}))
	assert.True(t, br.ContainsXY(Coord2{1, 0}))
	assert.True(t, br.ContainsXY(Coord2{2, 0}))
	assert.False(t, br.ContainsXY(Coord2{3, 0}))

	assert.False(t, br.ContainsXY(Coord2{-1, 1}))
	assert.False(t, br.ContainsXY(Coord2{0, 1}))
	assert.False(t, br.ContainsXY(Coord2{1, 1}))
	assert.False(t, br.ContainsXY(Coord2{2, 1}))
	assert.False(t, br.ContainsXY(Coord2{3, 1}))
}

func TestBricksSortedByColumn(t *testing.T) {
	bricks := []Brick{
		{Name: "A", Origin: Coord3{1, 0, 1}, Size: Coord3{1, 3, 1}},
		{Name: "B", Origin: Coord3{0, 0, 2}, Size: Coord3{3, 1, 1}},
		{Name: "C", Origin: Coord3{0, 2, 3}, Size: Coord3{3, 1, 1}},
		{Name: "D", Origin: Coord3{0, 0, 4}, Size: Coord3{1, 3, 1}},
		{Name: "E", Origin: Coord3{2, 0, 5}, Size: Coord3{1, 3, 1}},
		{Name: "F", Origin: Coord3{0, 1, 6}, Size: Coord3{3, 1, 1}},
		{Name: "G", Origin: Coord3{1, 1, 8}, Size: Coord3{1, 1, 2}},
	}

	byCol := BricksSortedByColumns(bricks)
	for col, brickIDs := range byCol {
		fmt.Println()
		for _, brickID := range brickIDs {
			fmt.Printf("%#v: %v\n", col, bricks[brickID])
		}
	}

	assert.Len(t, byCol[Coord2{0, 0}], 2)
	assert.Len(t, byCol[Coord2{1, 0}], 2)
	assert.Len(t, byCol[Coord2{2, 0}], 2)
	assert.Len(t, byCol[Coord2{0, 1}], 2)
	assert.Len(t, byCol[Coord2{1, 1}], 3)
	assert.Len(t, byCol[Coord2{2, 1}], 2)
	assert.Len(t, byCol[Coord2{0, 2}], 2, byCol[Coord2{0, 2}])
	assert.Len(t, byCol[Coord2{1, 2}], 2)
	assert.Len(t, byCol[Coord2{2, 2}], 2)
}
