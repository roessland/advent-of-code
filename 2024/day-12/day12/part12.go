package day12

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/disjointset"
)

//go:embed input*.txt
var Input embed.FS

func GetID(x, y, width int) int {
	return y*width + x
}

// Funfact: The number of vertical edges is equal to the number of horizontal edges,
// so this is enough.
func IsVerticalEdge(m [][]byte, W, H, i, j, comp int, comps *disjointset.DisjointSet) bool {
	inComp := func(i, j int) int {
		if comps.Connected(GetID(j, i, W), comp) {
			return 1
		}
		return 0
	}

	dx := func(i, j int) int {
		return inComp(i, j+1) - inComp(i, j)
	}

	return dx(i, j) != 0 && dx(i+1, j) != dx(i, j)
}

func Part12(inputName string) (int, int) {
	m := aocutil.FSReadLinesAsBytes(Input, inputName)
	m = aocutil.PadMap(m, 2, '.')
	height := len(m)
	width := len(m[0])

	getId := func(x, y int) int {
		return GetID(x, y, width)
	}

	// Step 1: Find connected components
	components := disjointset.Make(height * width)
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]
			id := getId(x, y)
			if c == '.' {
				continue
			}

			if m[y-1][x] == c {
				components.Union(id, getId(x, y-1))
			}
			if m[y+1][x] == c {
				components.Union(id, getId(x, y+1))
			}
			if m[y][x-1] == c {
				components.Union(id, getId(x-1, y))
			}
			if m[y][x+1] == c {
				components.Union(id, getId(x+1, y))
			}
		}
	}

	// Area and perimeter for each component
	areas := map[int]int{}
	perimeters := map[int]int{}

	// Keep track of the extent of each component.
	// Just an optimization, not necessary for the solution
	// Takes runtime from 200 ms to 5 ms.
	yMin := map[int]int{}
	yMax := map[int]int{}
	xMin := map[int]int{}
	xMax := map[int]int{}

	// Part 1: Calculate area and perimeter for each component
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := m[y][x]
			if c == '.' {
				continue
			}

			comp := components.Find(getId(x, y))
			areas[comp]++

			if _, ok := yMin[comp]; !ok || y <= yMin[comp] {
				yMin[comp] = y
			}
			if _, ok := yMax[comp]; !ok || y >= yMax[comp] {
				yMax[comp] = y
			}
			if _, ok := xMin[comp]; !ok || x <= xMin[comp] {
				xMin[comp] = x
			}
			if _, ok := xMax[comp]; !ok || x >= xMax[comp] {
				xMax[comp] = x
			}

			if m[y-1][x] != c {
				perimeters[comp]++
			}
			if m[y+1][x] != c {
				perimeters[comp]++
			}
			if m[y][x-1] != c {
				perimeters[comp]++
			}
			if m[y][x+1] != c {
				perimeters[comp]++
			}
		}
	}

	sum1 := 0
	sum2 := 0
	for c := range areas {
		sum1 += areas[c] * perimeters[c]
	}

	// Part 2: Calculate number of sides for each component
	for comp := range areas {
		sides := 0
		for y := yMin[comp] - 1; y < yMax[comp]+1; y++ {
			for x := xMin[comp] - 1; x < xMax[comp]+1; x++ {
				if IsVerticalEdge(m, width, height, y, x, comp, components) {
					sides += 2
				}
			}
		}
		sum2 += areas[comp] * sides
	}

	return sum1, sum2
}
