package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/gopkg/priorityqueue"
	"log"
	"math"
	"os"
)

const Height = 7
const Width = 13

var targetRoomAtJ = map[int]byte{
	3: 'A',
	5: 'B',
	7: 'C',
	9: 'D',
}

type Burrow [Height][Width]byte

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
	Pos    Pos
	Energy float64
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
func (b Burrow) IsWinningState() bool {
	//#############
	//#...........#
	//###C#C#A#B###
	//  #D#D#B#A#
	//  #########
	for col, amph := 3, byte('A'); amph <= 'D'; col, amph = col+2, amph+1 {
		for row := 2; row < 4; row++ {
			if b[row][col] != amph {
				return false
			}
		}
	}
	return true
}

func (b Burrow) AmphipodPositions() []Pos {
	var poses []Pos
	for i := 1; i < Height-1; i++ {
		for j := 1; j < Width-1; j++ {
			pos := Pos{i, j}
			r := b.At(pos)
			if 'A' <= r && r <= 'Z' {
				poses = append(poses, pos)
			}
		}
	}
	return poses
}

func ReadInput() (burrow Burrow) {
	f, err := os.Open("input_part1.txt")
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

func Dijkstra(b0 Burrow) {
	energyTo := map[Burrow]float64{b0: 0.0}
	Q := priorityqueue.New[Burrow]()
	Q.Push(b0, 0)

	for Q.Len() > 0 {
		b := Q.Pop()
		if b.IsWinningState() {
			b.Print()
			fmt.Println(energyTo[b])
			return
		}
		for _, amphipodPos0 := range b.AmphipodPositions() {
			for _, move := range GetMoves(b, amphipodPos0) {
				bNext := DoMove(b, amphipodPos0, move.Pos)
				currEnergy, ok := energyTo[bNext]
				if !ok {
					currEnergy = math.MaxFloat64
				}

				altEnergy := energyTo[b] + move.Energy
				if altEnergy < currEnergy {
					energyTo[bNext] = altEnergy
					Q.Push(bNext, altEnergy)
				}
			}
		}
	}
	fmt.Println(GetMoves(b0, Pos{2, 9}))
}

func DoMove(b Burrow, p0, p1 Pos) Burrow {
	b[p1.I][p1.J] = b.At(p0)
	b[p0.I][p0.J] = '.'
	return b
}

func GetMoves(b Burrow, p0 Pos) []Move {
	energyPerStep := EnergyPerStep(b.At(p0))
	typ := b.At(p0)

	// Moving out of a completed room is forbidden
	if targetRoomAtJ[p0.J] == typ && b.At(p0.Down()) == typ {
		return nil
	}

	visited := map[Pos]Move{}

	var dfs func(p Pos, energy float64)
	dfs = func(p Pos, energy float64) {
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

func EnergyPerStep(amph byte) float64 {
	var baseCost float64
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
	Dijkstra(b)
}
