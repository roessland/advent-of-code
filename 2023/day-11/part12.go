package main

import (
	"fmt"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

type Galaxy struct {
	col0, row0 int
}

type Dist struct {
	fixed, expanding int
}

func (d Dist) Sum() int {
	return d.fixed + d.expanding
}

func (dist Dist) Expand(factor int) Dist {
	return Dist{
		fixed:     dist.fixed,
		expanding: dist.expanding * factor,
	}
}

func main() {
	input := ReadInput()

	galaxies := make([]Galaxy, 0)
	rowHasGalaxy := make([]bool, len(input))
	colHasGalaxy := make([]bool, len(input[0]))
	for iy, row := range input {
		for ix, col := range row {
			if col == '#' {
				galaxies = append(galaxies, Galaxy{col0: ix, row0: iy})
				rowHasGalaxy[iy] = true
				colHasGalaxy[ix] = true
			}
		}
	}

	distsSum := Dist{fixed: 0, expanding: 0}
	for ia, ga := range galaxies {
		for ib, gb := range galaxies {
			if ib >= ia {
				continue
			}
			d := GetDist(ga, gb, rowHasGalaxy, colHasGalaxy)
			distsSum.fixed += d.fixed
			distsSum.expanding += d.expanding
		}
	}

	fmt.Println("Part 1:", distsSum.Expand(2).Sum())
	fmt.Println("Part 2:", distsSum.Expand(1000000).Sum())
}

// Split the distance into two parts: fixed and expanding.
func GetDist(ga, gb Galaxy, rowHasGalaxy, colHasGalaxy []bool) Dist {
	d := Dist{fixed: 0, expanding: 0}
	if ga.col0 > gb.col0 {
		ga, gb = gb, ga
	}
	for c := ga.col0; c < gb.col0; c++ {
		if colHasGalaxy[c] {
			d.fixed++
		} else {
			d.expanding++
		}
	}

	if ga.row0 > gb.row0 {
		ga, gb = gb, ga
	}
	for r := ga.row0; r < gb.row0; r++ {
		if rowHasGalaxy[r] {
			d.fixed++
		} else {
			d.expanding++
		}
	}

	return d
}

func ReadInput() []string {
	return aocutil.ReadLines("input.txt")
}
