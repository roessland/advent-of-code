// Package day04 solves AoC 2025 Day 4
package day04

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var InputFile embed.FS

func ReadInput(inputName string) (grid [][]byte) {
	return aocutil.ReadLinesAsBytes(inputName)
}

type Pos struct {
	I, J int
}

func AdjacentPositions(grid [][]byte, p Pos) (ns []Pos) {
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			c := Pos{p.I + di, p.J + dj}
			if c.I >= 0 && c.I < len(grid) && c.J >= 0 && c.J < len(grid[0]) {
				ns = append(ns, c)
			}
		}
	}
	return ns
}

// NumRollsOfPaper counts the '@' in the given positions
func NumRollsOfPaper(grid [][]byte, ps []Pos) (num int) {
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
	adjacentPositions := AdjacentPositions(grid, p)
	return NumRollsOfPaper(grid, adjacentPositions) < 4
}

func ForEachPosition(grid [][]byte, fn func(p Pos, c byte)) {
	for i := range grid {
		for j := range grid[i] {
			fn(Pos{i, j}, grid[i][j])
		}
	}
}

// Part1 returns number of rolls of paper that can be accessed by a forklift
func Part1(grid [][]byte) (accessible int) {
	ForEachPosition(grid, func(p Pos, c byte) {
		if Accessible(grid, p) {
			accessible++
		}
	})
	return accessible
}

// Part2 returns number of rolls of paper that can eventually be accessed
func Part2(grid [][]byte) (accessible int) {
	for {
		found := false
		ForEachPosition(grid, func(p Pos, c byte) {
			if Accessible(grid, p) {
				accessible++
				grid[p.I][p.J] = '.'
				found = true
			}
		})
		if !found {
			return accessible
		}
	}
}
