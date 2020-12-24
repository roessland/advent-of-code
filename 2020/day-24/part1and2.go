package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Split(str string) []Dir {
	str = strings.ReplaceAll(str, "se", "sE ")
	str = strings.ReplaceAll(str, "sw", "sW ")
	str = strings.ReplaceAll(str, "ne", "nE ")
	str = strings.ReplaceAll(str, "nw", "nW ")
	str = strings.ReplaceAll(str, "e", "e ")
	str = strings.ReplaceAll(str, "w", "w ")
	str = strings.ReplaceAll(str, "W", "w")
	str = strings.ReplaceAll(str, "E", "e")
	str = strings.Trim(str, " ")
	var dirs []Dir
	for _, s := range strings.Split(str, " ") {
		dirs = append(dirs, NewDir(s))
	}
	return dirs
}

type Dir struct {
	E, NE, SE int
}

var E = Dir{1,0,0}
var NE = Dir{0, 1, 0}
var SE = Dir{0, 0, 1}
var W = Dir{-1,0,0}
var SW = Dir{0,-1,0}
var NW = Dir{0,0,-1}


func (u Dir) Mul(n int) Dir {
	return Dir{u.E * n, u.NE * n, u.SE * n}
}

// Canonicalize makes the SE component 0.
// (0, 0, 0) = (1, -1, 0) - (0, 0, 1) = (1, -1, -1)
func (u Dir) Canonicalize() Dir {
		return u.Add(Dir{1, -1, -1}.Mul(u.SE))
}

func NewDir(dir string) Dir {
	switch dir {
	case "e":
		return Dir{1, 0, 0}
	case "ne":
		return Dir{0, 1, 0}
	case "nw":
		return Dir{0, 0, -1}
	case "w":
		return Dir{-1, 0, 0}
	case "sw":
		return Dir{0, -1, 0}
	case "se":
		return Dir{0, 0, 1}
	default:
		panic("nope")
	}
}

func (u Dir) Add(v Dir) Dir {
	return Dir{E: u.E + v.E, NE: u.NE + v.NE, SE: u.SE + v.SE}
}

func Sum(dirs []Dir) Dir {
	here := Dir{}
	for _, dir := range dirs {
		here = here.Add(dir)
	}
	return here
}

type Color bool
const Black = true
const White = false

type Grid map[Dir]Color // black is true

func (g Grid) Copy() Grid {
	c := make(Grid)
	for k, v := range g {
		c[k] = v
	}
	return c
}

func (g Grid) Flip(pos Dir) {
	g[pos.Canonicalize()] = !g[pos.Canonicalize()]
}

func (g Grid) CountBlack() int {
	blackCount := 0
	for _, black := range g {
		if black {
			blackCount++
		}
	}
	return blackCount
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	grid := make(Grid)
	for scanner.Scan() {
		line := scanner.Text()
		pos := Sum(Split(line)).Canonicalize()
		grid[pos] = !grid[pos]
	}
	fmt.Println("Part 1:", grid.CountBlack())

	for i := 0; i < 100; i++ {
		blackNeighborCount := make(map[Dir]int)
		for pos, color := range grid {
			if color == White {
				continue
			}
			blackNeighborCount[pos] += 0 // Important. So we can iterate this pos in the blackNeighborCount loop.
			for _, dir := range []Dir{E, W, NE, SW, SE, NW} {
				blackNeighborCount[pos.Add(dir).Canonicalize()]++
			}
		}
		for pos, count := range blackNeighborCount {
			if grid[pos] == Black && (count == 0 || count > 2) {
				grid.Flip(pos)
			} else if grid[pos] == White && count == 2 {
				grid.Flip(pos)
			}
		}
	}
	fmt.Println("Part 2:", grid.CountBlack())
}
