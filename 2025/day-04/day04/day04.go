// Package day04 solves AoC 2025 Day 4
package day04

import (
	"embed"
	"iter"

	"github.com/roessland/advent-of-code/2025/aocutil"
)

//go:embed input*.txt
var InputFS embed.FS

func ReadInput(inputName string) [][]byte {
	return aocutil.FSReadLinesAsBytes(InputFS, inputName)
}

type Pos struct {
	I, J int
}

func (p Pos) Valid(grid [][]byte) bool {
	return p.I >= 0 && p.I < len(grid) && p.J >= 0 && p.J < len(grid[0])
}

func AdjacentPositions(grid [][]byte, p Pos) []Pos {
	ns := make([]Pos, 0, 8)
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			c := Pos{p.I + di, p.J + dj}
			if c.Valid(grid) {
				ns = append(ns, c)
			}
		}
	}
	return ns
}

// NumRollsOfPaper counts the '@' in the given positions
func NumRollsOfPaper(grid [][]byte, ps []Pos) int {
	num := 0
	for _, p := range ps {
		if grid[p.I][p.J] == '@' {
			num++
		}
	}
	return num
}

// Accessible returns whether a position is paper and accessible by forklift
func Accessible(grid [][]byte, p Pos) bool {
	if grid[p.I][p.J] != '@' {
		return false
	}
	return NumRollsOfPaper(grid, AdjacentPositions(grid, p)) < 4
}

func Positions(grid [][]byte) iter.Seq[Pos] {
	return func(yield func(Pos) bool) {
		for i := range grid {
			for j := range grid[i] {
				if !yield(Pos{i, j}) {
					return
				}
			}
		}
	}
}

// Part1 returns number of rolls of paper that can be accessed by a forklift
func Part1(grid [][]byte) int {
	accessible := 0
	for p := range Positions(grid) {
		if Accessible(grid, p) {
			accessible++
		}
	}
	return accessible
}

// Part2 returns number of rolls of paper that can eventually be accessed
func Part2(grid [][]byte) int {
	accessible := 0
	for {
		found := false
		for p := range Positions(grid) {
			if Accessible(grid, p) {
				accessible++
				grid[p.I][p.J] = '.'
				found = true
			}
		}
		if !found {
			return accessible
		}
	}
}
