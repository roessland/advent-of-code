package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"github.com/roessland/gopkg/priorityqueue"
	"log"
	"math"
	"os"
	"time"
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

func (b *Burrow) At(pos Pos) byte {
	return b[pos.I][pos.J]
}

func (b *Burrow) Print() {
	for _, line := range b {
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

func (b *Burrow) IsWinningState() bool {
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

func (b *Burrow) AmphipodPositions() []Pos {
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
	f, err := os.Open("input_part2.txt")
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

func H(b Burrow) float64 {
	estimate := 0.0
	for _, amphipodPos := range b.AmphipodPositions() {
		typ := b.At(amphipodPos)
		energyPerStep := EnergyPerStep(typ)
		currentCol := amphipodPos.J
		targetCol := map[byte]int{
			'A': 3,
			'B': 5,
			'C': 7,
			'D': 9,
		}[typ]

		steps := mathutil.AbsInt(currentCol - targetCol)
		if currentCol != targetCol {
			steps += 2 * (amphipodPos.I - 1)
		}

		estimate += float64(steps) * energyPerStep
	}
	return estimate
}

func AStar(b0 Burrow) (cost float64) {
	energyTo := map[Burrow]float64{b0: 0.0}
	Q := priorityqueue.New[Burrow]()
	Q.Push(b0, H(b0))

	for Q.Len() > 0 {
		b := Q.Pop()
		if b.IsWinningState() {
			bWinning := b
			cost = energyTo[bWinning]
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
					Q.Push(bNext, altEnergy+H(bNext))
				}
			}
		}
	}
	return math.MaxFloat64
}

func DoMove(b Burrow, p0, p1 Pos) Burrow {
	b[p1.I][p1.J] = b.At(p0)
	b[p0.I][p0.J] = '.'
	return b
}

func IsRoomCompleted(b Burrow, j int) bool {
	return b[2][j] != '#' && b[2][j] == b[3][j] && b[3][j] == b[4][j] && b[4][j] == b[5][j]
}

func RoomContainsOtherTypes(b Burrow, j int, typ byte) bool {
	for i := 2; i < Height-1; i++ {
		if b[i][j] != '.' && b[i][j] != typ {
			return true
		}
	}
	return false
}

func IsRoomCol(j int) bool {
	return j == 3 || j == 5 || j == 7 || j == 9
}

func GetMoves(b Burrow, p0 Pos) []Move {
	energyPerStep := EnergyPerStep(b.At(p0))
	typ := b.At(p0)

	// Moving out of a completed room is forbidden
	if IsRoomCompleted(b, p0.J) {
		return nil
	}

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
		if p.I != p0.I && p.I > 1 && targetRoomAtJ[p.J] != typ {
			return false
		}
		// Forbidden to move into a room that contains amphipods of the wrong type
		if p.I > 1 && RoomContainsOtherTypes(b, p.J, typ) {
			return false
		}
		// Not going all the way in to a room is forbidden
		if b.At(p.Down()) == '.' {
			return false
		}
		// Standing outside a room is forbidden
		if p.I == 1 && IsRoomCol(p.J) {
			return false
		}
		// Standing above someone of another type is forbidden
		if 'A' <= b.At(p.Down()) && b.At(p.Down()) <= 'D' && b.At(p.Down()) != typ {
			return false
		}
		return true
	}

	visited := map[Pos]bool{}
	var moves []Move

	var dfs func(p Pos, energy float64)
	dfs = func(p Pos, energy float64) {
		if b.At(p) != '.' && p != p0 {
			return
		}
		if visited[p] {
			return
		} else {
			visited[p] = true
			if isLegalMove(p) {
				moves = append(moves, Move{p, energy})
			}
		}

		dfs(p.Up(), energy+energyPerStep)
		dfs(p.Down(), energy+energyPerStep)
		dfs(p.Left(), energy+energyPerStep)
		dfs(p.Right(), energy+energyPerStep)
	}
	dfs(p0, 0)
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
	b0 := ReadInput()
	t0 := time.Now()
	cost := AStar(b0)
	fmt.Println(cost)
	fmt.Println(time.Since(t0))
}
