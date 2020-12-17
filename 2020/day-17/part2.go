package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	X, Y, Z, W int
}

func (p Pos) Neighbors() []Pos {
	var neighbors []Pos
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					neighbors = append(neighbors, Pos{p.X+x, p.Y+y, p.Z+z, p.W+w})
				}
			}
		}
	}
	return neighbors
}

type State map[Pos]bool

func Next(prevActive State) (active State) {
	neighbors := make(map[Pos]int)
	active = make(State)
	for pos, _ := range prevActive {
		for _, nPos := range pos.Neighbors() {
			neighbors[nPos]++
		}
	}
	for pos, count := range neighbors {
		if prevActive[pos] {
			if count == 2 || count == 3 {
				active[pos] = true
			}
		} else {
			if count == 3 {
				active[pos] = true
			}
		}
	}

	return active
}

func main() {
	active := make(State)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			if c == '#' {
				active[Pos{i, j, 0, 0}] = true
			}
		}
		i++
	}

	for i := 0; i < 6; i++ {
		active = Next(active)
	}

	fmt.Println(len(active))
}
