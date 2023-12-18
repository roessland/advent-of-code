package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	fmt.Println("Part 1:", compute(0))
	fmt.Println("Part 2:", compute(1))
}

func compute(smudges int) int {
	patterns := ReadInput("input.txt")
	sum := 0
	for _, pattern := range patterns {
		j, hasVertical := FindRowMirror(pattern, smudges)
		if hasVertical {
			sum += (j + 1) * 100
		}

		i, hasHorizontal := FindRowMirror(pattern.Transposed(), smudges)
		if hasHorizontal {
			sum += (i + 1)
		}
	}
	return sum
}

// FindRowMirror finds the location where a mirror is located,
// given the exact number of smudges.
func FindRowMirror(p Pattern, smudges int) (loc int, hasMirror bool) {
	for loc = 0; loc < p.numRows-1; loc++ {
		totalDist := 0
		for i, j := loc, loc+1; ; i, j = i-1, j+1 {
			a, b := p.Row(i), p.Row(j)

			if a == "" && b == "" {
				break
			}

			if a == "" || b == "" {
				continue
			}

			totalDist += Dist(a, b)

			if totalDist > smudges {
				break
			}
		}
		if totalDist == smudges {
			return loc, true
		}
	}
	return -1, false
}

func Dist(a, b string) int {
	numDiff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			numDiff++
		}
	}
	return numDiff
}

// Pattern from input. Has the same data in row-major and col-major order for
// easy row/col access.
type Pattern struct {
	data    []byte
	dataT   []byte
	numRows int
	numCols int
}

func (p Pattern) Transposed() Pattern {
	return Pattern{
		data:    p.dataT,
		dataT:   p.data,
		numRows: p.numCols,
		numCols: p.numRows,
	}
}

func (p Pattern) Row(y int) string {
	if y < 0 || y >= p.numRows {
		return ""
	}
	return string(p.data[y*p.numCols : (y+1)*p.numCols])
}

func (p Pattern) Col(x int) string {
	if x < 0 || x >= p.numCols {
		return ""
	}
	return string(p.dataT[x*p.numRows : (x+1)*p.numRows])
}

// Reflect reflects an element around the mirror after i and before i+1.
// For example, Reflect(0, 1) = 3.
//
//	x  |  x
//	0 1 2 3 4 5
func Reflect(i, around int) int {
	return 2*around - i + 1
}

func ReadInput(filename string) []Pattern {
	f := aocutil.ReadFile(filename)
	patterns := []Pattern{}
	for _, patternStr := range strings.Split(f, "\n\n") {
		lines := strings.Split(patternStr, "\n")
		pattern := Pattern{
			numRows: len(lines),
			numCols: len(lines[0]),
			data:    make([]byte, len(lines)*len(lines[0])),
			dataT:   make([]byte, len(lines)*len(lines[0])),
		}
		for y, line := range lines {
			for x, char := range line {
				pattern.data[y*pattern.numCols+x] = byte(char)
				pattern.dataT[x*pattern.numRows+y] = byte(char)
			}
		}
		patterns = append(patterns, pattern)
	}
	return patterns
}
