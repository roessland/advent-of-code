package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var targetRoomAtJ = map[int]byte{
	3: 'A',
	5: 'B',
	7: 'C',
	9: 'D',
}

type Burrow [5][13]byte

type Pos struct {
	I, J int
}

func (p Pos) Up() Pos {
	return Pos{p.I - 1, p.J}
}

func (p Pos) Down() Pos {
	return Pos{p.I + 1, p.J}
}

func (p Pos) Left() Pos {
	return Pos{p.I, p.J - 1}
}

func (p Pos) Right() Pos {
	return Pos{p.I, p.J + 1}
}

type Move struct {
	Pos  Pos
	Dist int
}

func (b Burrow) At(pos Pos) byte {
	return b[pos.I][pos.J]
}

func (b Burrow) Print() {
	for _, line := range b {
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

func ReadInput() (burrow Burrow) {
	f, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		copy(burrow[i][:], line)
		i++
	}
	return burrow
}

func BFS(b Burrow) {
	fmt.Println(GetMoves(b, Pos{2, 9}))
}

func GetMoves(b Burrow, p0 Pos) []Move {
	energyPerStep := EnergyPerStep(b.At(p0))
	typ := b.At(p0)

	// Moving out of a completed room is forbidden
	if targetRoomAtJ[p0.J] == typ && b.At(p0.Down()) == typ {
		return nil
	}

	visited := map[Pos]Move{}

	var dfs func(p Pos, energy int)
	dfs = func(p Pos, energy int) {
		if b.At(p) != '.' && p != p0 {
			return
		}
		_, alreadyVisited := visited[p]
		if alreadyVisited {
			return
		} else {
			visited[p] = Move{p, energy}
		}
		dfs(p.Up(), energy+energyPerStep)
		dfs(p.Down(), energy+energyPerStep)
		dfs(p.Left(), energy+energyPerStep)
		dfs(p.Right(), energy+energyPerStep)
	}
	dfs(p0, 0)

	isLegalMove := func(p Pos) bool {
		if p == p0 {
			return false
		}
		// Not moving vertically is forbidden.
		if p.I == p0.I {
			return false
		}
		// Not moving horizontally is forbidden.
		if p.J == p0.J {
			return false
		}
		// Forbidden to move into wrong room
		if p.I > p0.I && targetRoomAtJ[p.J] != typ {
			return false
		}
		// Standing outside room, or not going all the way
		// inside a room is forbidden.
		if b.At(p.Down()) == '.' {
			return false
		}
		// Standing above someone of another type is forbidden
		if 'A' <= b.At(p.Down()) && b.At(p.Down()) <= 'D' && b.At(p.Down()) != typ {
			return false
		}
		return true
	}

	var moves []Move
	for _, move := range visited {
		if isLegalMove(move.Pos) {
			moves = append(moves, move)
		}
	}
	return moves
}

func EnergyPerStep(amph byte) int {
	var baseCost int
	switch amph {
	case 'A':
		baseCost = 1
	case 'B':
		baseCost = 10
	case 'C':
		baseCost = 100
	case 'D':
		baseCost = 1000
	default:
		panic("illegal amphipod type")
	}
	return baseCost
}

func main() {
	b := ReadInput()
	BFS(b)
}
