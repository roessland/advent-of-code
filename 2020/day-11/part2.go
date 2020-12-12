package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tile int

const Floor Tile = 0
const EmptySeat Tile = 1
const OccupiedSeat Tile = 2

func NewTile(c rune) Tile {
	switch c {
	case 'L':
		return EmptySeat
	case '#':
		return OccupiedSeat
	case '.':
		return Floor
	default:
		panic("no such tile " + string(c))
	}
}

type Pos struct {
	I, J int
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.I + v.I, p.J + v.J}
}

func (a Area) RayTrace(p0 Pos, v Pos) Tile {
	d := 1
	for {
		p := p0.Add(Pos{d*v.I, d*v.J})
		if p.I < 0 || a.Height <= p.I || p.J < 0 || a.Width <= p.J {
			return Floor
		}
		if a.Tiles[p] == EmptySeat || a.Tiles[p] == OccupiedSeat {
			return a.Tiles[p]
		}
		d++
	}
}

func (a Area) CountAdjacent(p Pos) map[Tile]int {
	freqs := make(map[Tile]int)
	freqs[a.RayTrace(p, Pos{1,1})]++
	freqs[a.RayTrace(p, Pos{1,0})]++
	freqs[a.RayTrace(p, Pos{1,-1})]++
	freqs[a.RayTrace(p, Pos{-1,1})]++
	freqs[a.RayTrace(p, Pos{-1,0})]++
	freqs[a.RayTrace(p, Pos{-1,-1})]++
	freqs[a.RayTrace(p, Pos{0,1})]++
	freqs[a.RayTrace(p, Pos{0,-1})]++
	return freqs
}

type Area struct {
	Tiles map[Pos]Tile
	Height int
	Width int
}

func (a Area) NextTile(p Pos) Tile {
	if a.Tiles[p] == Floor {
		return Floor
	}
	adjacent := a.CountAdjacent(p)
	// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	if a.Tiles[p] == EmptySeat && adjacent[OccupiedSeat] == 0 {
		return OccupiedSeat
	}
	// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
	if a.Tiles[p] == OccupiedSeat && adjacent[OccupiedSeat] >= 5 {
		return EmptySeat
	}
	// Otherwise, the seat's state does not change.
	return a.Tiles[p]
}

func (a0 Area) Evolve() (a1 Area, changed bool) {
	a1 = Area{make(map[Pos]Tile), a0.Height, a0.Width}
	for p, t0 := range a0.Tiles {
		a1.Tiles[p] = a0.NextTile(p)
		if a1.Tiles[p] != t0 {
			changed = true
		}
	}
	return
}

func (a0 Area) EvolveToSteadyState() Area {
	an := a0
	changed := false
	for {
		an, changed = an.Evolve()
		// fmt.Println(an.CountOccupied())
		if !changed {
			return an
		}
	}
}

func (a Area) CountOccupied() (numOccupied int) {
	for _, t := range a.Tiles {
		if t == OccupiedSeat {
			numOccupied++
		}
	}
	return numOccupied
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	area := Area{Tiles: make(map[Pos]Tile)}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		if i+1 > area.Height {
			area.Height = i+1
		}
		for j, c := range scanner.Text() {
			if j+1 > area.Width {
				area.Width = j+1
			}
			area.Tiles[Pos{i,j}] = NewTile(c)
		}
		i++
	}

	fmt.Println(area.EvolveToSteadyState().CountOccupied())
}