package day04

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type M [][]byte

func (m M) At(x, y int) byte {
	if x < 0 || x >= len(m) || y < 0 || y >= len(m[x]) {
		return '.'
	}
	return m[y][x]
}

type Direction struct {
	Dx, Dy int
}

var directions = []Direction{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}

func (m M) Ats(x, y int, d Direction, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += string(m.At(x+d.Dx*i, y+d.Dy*i))
	}
	return s
}

func (m M) IsXmas(x, y int) int {
	down := m.Ats(x-1, y-1, Direction{1, 1}, 3)
	if down != "MAS" && down != "SAM" {
		return 0
	}

	up := m.Ats(x-1, y+1, Direction{1, -1}, 3)
	if up != "MAS" && up != "SAM" {
		return 0
	}

	return 1
}

func ReadInput(inputName string) M {
	return M(aocutil.ReadLinesAsBytes(inputName))
}

func (m M) CountStartingAt(x, y int) int {
	n := 0
	for _, d := range directions {
		if m.Ats(x, y, d, 4) == "XMAS" {
			n++
		}
	}
	return n
}

func Part12(inputName string) (int, int) {
	m := ReadInput(inputName)
	total1 := 0
	total2 := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m.At(x, y) == 'X' {
				total1 += m.CountStartingAt(x, y)
			}
			if m.At(x, y) == 'A' {
				total2 += m.IsXmas(x, y)
			}
		}
	}
	return total1, total2
}
