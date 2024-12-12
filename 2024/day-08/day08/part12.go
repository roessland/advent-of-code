package day08

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Vec struct {
	X, Y int
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec) Sub(v2 Vec) Vec {
	return Vec{v.X - v2.X, v.Y - v2.Y}
}

func FindAntennas(input [][]byte) map[byte][]Vec {
	antennas := map[byte][]Vec{}
	for y, row := range input {
		for x, c := range row {
			if c != '.' {
				antennas[c] = append(antennas[c], Vec{x, y})
			}
		}
	}
	return antennas
}

func FindAntiNodes1(antennas []Vec, width, height int) map[Vec]bool {
	antiNodes := map[Vec]bool{}
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			a, b := antennas[i], antennas[j]

			ab := b.Sub(a)
			p1 := a.Add(ab).Add(ab)
			if WithinBounds(p1, width, height) {
				antiNodes[p1] = true
			}

			ba := a.Sub(b)
			p2 := b.Add(ba).Add(ba)
			if WithinBounds(p2, width, height) {
				antiNodes[p2] = true
			}
		}
	}
	return antiNodes
}

func FindAntiNodes2(antennas []Vec, width, height int) map[Vec]bool {
	antiNodes := map[Vec]bool{}
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			a, b := antennas[i], antennas[j]

			ab := b.Sub(a)
			p1 := b
			for WithinBounds(p1, width, height) {
				antiNodes[p1] = true
				p1 = p1.Add(ab)
			}

			ba := a.Sub(b)
			p2 := a
			for WithinBounds(p2, width, height) {
				antiNodes[p2] = true
				p2 = p2.Add(ba)
			}
		}
	}
	return antiNodes
}

func WithinBounds(v Vec, width, height int) bool {
	return v.X >= 0 && v.X < width && v.Y >= 0 && v.Y < height
}

func Part12(inputName string) (int, int) {
	input := aocutil.FSReadLinesAsBytes(Input, inputName)
	height := len(input)
	width := len(input[0])

	hasAntiNode1 := map[Vec]bool{}
	hasAntiNode2 := map[Vec]bool{}

	antennas := FindAntennas(input)
	for freq := range antennas {
		for loc := range FindAntiNodes1(antennas[freq], width, height) {
			hasAntiNode1[loc] = true
		}
		for loc := range FindAntiNodes2(antennas[freq], width, height) {
			hasAntiNode2[loc] = true
		}
	}

	sum1, sum2 := len(hasAntiNode1), len(hasAntiNode2)
	return sum1, sum2
}
