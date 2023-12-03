package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	schem := ReadInput()
	part1(schem)
	part2(schem)
}

func part1(schem Schematic) {
	sum := 0
	for y := 0; y < schem.Height(); y++ {
		for x := 0; x < schem.Width(); x++ {
			p := Pos{y, x}
			if schem.IsPartNumberRoot(p) {
				sum += schem.Value(p)
			}
		}
	}
	fmt.Println(sum)
}

func part2(schem Schematic) {
	// Assign every number to all adjacent gears.
	numsForGear := map[Pos][]int{}
	for y := 0; y < schem.Height(); y++ {
		for x := 0; x < schem.Width(); x++ {
			p := Pos{y, x}
			if !schem.IsPartNumberRoot(p) {
				continue
			}
			num := schem.Value(p)
			adjacentGears := schem.GetAdjacentGears(p)
			for _, gearPos := range adjacentGears {
				numsForGear[gearPos] = append(numsForGear[gearPos], num)
			}
		}
	}

	// Sum gear ratio for actual gears (those with 2 numbers)
	sum := 0
	for _, nums := range numsForGear {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}
	fmt.Println(sum)
}

type Schematic []string

func (s Schematic) Height() int {
	return len(s)
}

func (s Schematic) Width() int {
	return len(s[0])
}

// At returns char at position, or '.' if out of bounds.
func (s Schematic) At(p Pos) byte {
	x, y := p.X, p.Y
	if y >= s.Height() || x >= s.Width() || y < 0 || x < 0 {
		return '.'
	}
	return s[y][x]
}

type Pos struct {
	Y, X int
}

func (p Pos) Left() Pos {
	return Pos{p.Y, p.X - 1}
}

func (p Pos) Right() Pos {
	return Pos{p.Y, p.X + 1}
}

func (p Pos) Up() Pos {
	return Pos{p.Y - 1, p.X}
}

func (p Pos) Down() Pos {
	return Pos{p.Y + 1, p.X}
}

func (p Pos) UpLeft() Pos {
	return Pos{p.Y - 1, p.X - 1}
}

func (p Pos) UpRight() Pos {
	return Pos{p.Y - 1, p.X + 1}
}

func (p Pos) DownLeft() Pos {
	return Pos{p.Y + 1, p.X - 1}
}

func (p Pos) DownRight() Pos {
	return Pos{p.Y + 1, p.X + 1}
}

func (s Schematic) IsDigit(p Pos) bool {
	return unicode.IsDigit(rune(s.At(p)))
}

// IsNumberRoot is true for the first digit in a number, and false otherwise.
func (s Schematic) IsNumberRoot(numberRoot Pos) bool {
	return s.IsDigit(numberRoot) && !s.IsDigit(numberRoot.Left())
}

// Value returns the number represented by the number root.
func (s Schematic) Value(r Pos) int {
	if !s.IsNumberRoot(r) {
		panic("Not a number root")
	}
	p := r
	n := 0
	for s.IsDigit(p) {
		n = 10*n + int(s.At(p)-'0')
		p = p.Right()
	}
	return n
}

func (s Schematic) IsPartNumberRoot(p Pos) bool {
	return s.IsNumberRoot(p) && len(s.GetAdjacentSymbols(p)) > 0
}

func (s Schematic) GetAdjacentSymbols(r Pos) map[Pos]byte {
	adjacent := s.GetAdjacent(r)
	adjacentSymbols := map[Pos]byte{}
	for pos, c := range adjacent {
		if c != '.' && !unicode.IsDigit(rune(c)) {
			adjacentSymbols[pos] = c
		}
	}
	return adjacentSymbols
}

func (s Schematic) GetAdjacentGears(r Pos) []Pos {
	adjacentSymbols := s.GetAdjacentSymbols(r)
	gearPositions := []Pos{}
	for pos, c := range adjacentSymbols {
		if c == '*' {
			gearPositions = append(gearPositions, pos)
		}
	}
	return gearPositions
}

// GetAdjacent returns symbols adjacent to the number with the given root.
func (s Schematic) GetAdjacent(r Pos) map[Pos]byte {
	adjacent := map[Pos]byte{}
	adjacent[r.Left()] += s.At(r.Left())
	adjacent[r.UpLeft()] += s.At(r.UpLeft())
	adjacent[r.DownLeft()] += s.At(r.DownLeft())
	for s.IsDigit(r) {
		adjacent[r.Up()] += s.At(r.Up())
		adjacent[r.Down()] += s.At(r.Down())
		r = r.Right()
	}
	adjacent[r.Up()] += s.At(r.Up())
	adjacent[r] += s.At(r)
	adjacent[r.Down()] += s.At(r.Down())

	return adjacent
}

func ReadInput() Schematic {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	schem := make(Schematic, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		schem = append(schem, strings.TrimSpace(line))
	}

	return schem
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
