package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/gopkg/disjointset"
	"log"
	"os"
	"sort"
)



type Map struct {
	Width, Height int
	heights       []int
}

type Pos struct {
	I, J int
}

func (m Map) At(pos Pos) int {
	return m.heights[pos.J*m.Width+pos.I]
}

// Id returns an unique ID for each position.
// This is needed by the DisjointSet/UnionFind implementation.
func (m Map) Id(pos Pos) int {
	return pos.J*m.Width + pos.I
}

func (m Map) AllPositions() <-chan Pos {
	ch := make(chan Pos)
	go func() {
		for j := 0; j < m.Height; j++ {
			for i := 0; i < m.Width; i++ {
				ch <- Pos{i, j}
			}
		}
		close(ch)
	}()
	return ch
}

func (m Map) AdjacentPositions(pos Pos) []Pos {
	var legalAdjPositions []Pos
	for _, adjPos := range []Pos{
		{pos.I - 1, pos.J}, {pos.I + 1, pos.J},
		{pos.I, pos.J - 1}, {pos.I, pos.J + 1},
	} {
		if adjPos.I < 0 || adjPos.I >= m.Width {
			continue
		}
		if adjPos.J < 0 || adjPos.J >= m.Height {
			continue
		}
		legalAdjPositions = append(legalAdjPositions, adjPos)
	}
	return legalAdjPositions
}

func main() {
	m := ReadInput()
	part1(m)
	part2(m)
}

func part1(m Map) {
	totalRisk := 0
	for pos := range m.AllPositions() {
		height := m.At(pos)
		adjPositions := m.AdjacentPositions(pos)
		lowPoint := true
		for _, adjPosition := range adjPositions {
			adjHeight := m.At(adjPosition)
			if height >= adjHeight {
				lowPoint = false
				break
			}
		}
		if lowPoint {
			risk := 1 + height
			totalRisk += risk
		}
	}
	fmt.Println("Part 1:", totalRisk)
}

func part2(m Map) {
	// Start with every tile being its own component.
	basins := disjointset.Make(m.Height * m.Width)
	for pos := range m.AllPositions() {
		height := m.At(pos)
		if height == 9 {
			continue // Don't join walls with anything.
		}
		adjPositions := m.AdjacentPositions(pos)
		for _, adjPos := range adjPositions {
			adjHeight := m.At(adjPos)
			if adjHeight != 9 {
				// Same basin -- connect the components.
				basins.Union(m.Id(pos), m.Id(adjPos))
			}
		}
	}

	// Find the basin ID for every tile.
	basinSizes := make([]int, m.Width*m.Height)
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			pos := Pos{i, j}
			basinId := basins.Find(m.Id(pos))
			basinSizes[basinId]++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	fmt.Println("Part 2:", result)
}

func ReadInput() Map {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	m := Map{
		Height:  len(lines),
		Width:   len(lines[0]),
		heights: make([]int, len(lines)*len(lines[0])),
	}
	for j, line := range lines {
		for i, c := range line {
			m.heights[j*m.Width+i] = int(c - '0')
		}
	}
	return m
}
