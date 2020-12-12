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

func (a Area) CountAdjacent(p Pos) map[Tile]int {
	freqs := make(map[Tile]int)
	freqs[a[p.Add(Pos{1,1})]]++
	freqs[a[p.Add(Pos{1,0})]]++
	freqs[a[p.Add(Pos{1,-1})]]++
	freqs[a[p.Add(Pos{-1,1})]]++
	freqs[a[p.Add(Pos{-1,0})]]++
	freqs[a[p.Add(Pos{-1,-1})]]++
	freqs[a[p.Add(Pos{0,1})]]++
	freqs[a[p.Add(Pos{0,-1})]]++
	return freqs
}

type Area map[Pos]Tile

func (a Area) NextTile(p Pos) Tile {
	if a[p] == Floor {
		return Floor
	}
	adjacent := a.CountAdjacent(p)
	// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	if a[p] == EmptySeat && adjacent[OccupiedSeat] == 0 {
		return OccupiedSeat
	}
	// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
	if a[p] == OccupiedSeat && adjacent[OccupiedSeat] >= 4 {
		return EmptySeat
	}
	// Otherwise, the seat's state does not change.
	return a[p]
}

func (a0 Area) Evolve() (a1 Area, changed bool) {
	a1 = make(Area)
	for p, t0 := range a0 {
		a1[p] = a0.NextTile(p)
		if a1[p] != t0 {
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
	for _, t := range a {
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

	area := make(Area)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		for j, c := range scanner.Text() {
			area[Pos{i,j}] = NewTile(c)
		}
		i++
	}

	fmt.Println(area.EvolveToSteadyState().CountOccupied())
}