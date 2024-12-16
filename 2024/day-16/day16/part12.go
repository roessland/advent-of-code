package day16

import (
	"embed"
	"fmt"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/digraph"
)

//go:embed input*.txt
var Input embed.FS

type Pos struct {
	Y, X int
}

func ReadInput(inputName string) (m [][]byte, start Pos, end Pos) {
	m0 := aocutil.FSReadLinesAsBytes(Input, inputName)
	m = make([][]byte, len(m0))
	for y, row := range m0 {
		m[y] = make([]byte, len(row))
		for x, c := range row {
			if c == 'S' {
				start = Pos{y, x}
			}
			if c == 'E' {
				end = Pos{y, x}
			}
			if c == '#' {
				m[y][x] = '#'
			} else {
				m[y][x] = '.'
			}
		}
	}
	return
}

func PrintMap(m [][]byte) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

type Dir int

const (
	E Dir = 0
	N Dir = 1
	W Dir = 2
	S Dir = 3
)

func DirFromPos(p Pos) Dir {
	switch p {
	case Pos{0, 1}:
		return E
	case Pos{-1, 0}:
		return N
	case Pos{0, -1}:
		return W
	case Pos{1, 0}:
		return S
	}
	panic("nah")
}

func CW(d Dir) Dir {
	return (d + 1) % 4
}

func CCW(d Dir) Dir {
	return (d + 3) % 4
}

type State struct {
	Pos Pos
	Dir Dir
}

func (s State) ID(Nx int) int {
	return s.Pos.Y*4*Nx + s.Pos.X*4 + int(s.Dir)
}

func MakeGraph(m [][]byte) digraph.Graph {
	Ny := len(m)
	Nx := len(m[0])
	Nd := 4
	N := Nd * Nx * Ny

	g := digraph.Graph{
		Nodes: make([]digraph.Node, N),
	}

	for y, row := range m {
		for x, cell := range row {
			if cell == '#' {
				continue
			}

			for _, delta := range []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				pos := Pos{y, x}
				dir := DirFromPos(delta)
				thisID := State{pos, dir}.ID(Nx)

				// Move forward
				if m[y+delta.Y][x+delta.X] == '.' {
					nextID := State{Pos{Y: pos.Y + delta.Y, X: pos.X + delta.X}, dir}.ID(Nx)
					edge := digraph.Edge{To: nextID, Weight: 1.0}
					g.Nodes[thisID].Neighbors = append(g.Nodes[thisID].Neighbors, edge)
				}

				// Rotate in place
				for _, nextDir := range []Dir{CW(dir), CCW(dir)} {
					nextID := State{pos, nextDir}.ID(Nx)
					edge := digraph.Edge{To: nextID, Weight: 1000.0}
					g.Nodes[State{pos, dir}.ID(Nx)].Neighbors = append(g.Nodes[State{pos, dir}.ID(Nx)].Neighbors, edge)
				}
			}
		}
	}

	return g
}

// pt 2 29682 too high

func Part12(inputName string) (int, int) {
	m, startPos, endPos := ReadInput(inputName)
	Nx := len(m[0])
	// fmt.Println(startPos, endPos)
	// PrintMap(m)
	g := MakeGraph(m)

	sourceID := State{Pos: startPos, Dir: E}.ID(Nx)

	// fmt.Println("Dijkstraing")
	dist, prevs := digraph.DijkstraAll(g, sourceID)

	endID1 := State{Pos: endPos, Dir: E}.ID(Nx)
	endID2 := State{Pos: endPos, Dir: N}.ID(Nx)
	// fmt.Println(dist[endID1])
	// fmt.Println(dist[endID2])

	// Backtrack to find all nodes on the best path
	onBestPath := make(map[int]bool)
	var dfs func(int)
	dfs = func(id int) {
		if onBestPath[id] {
			return
		}
		onBestPath[id] = true
		for prev := range prevs[id] {
			dfs(prev)
		}
	}

	// Only backtrack from the nearest end position (there can be two end directions)
	part1 := 0
	if dist[endID1] < dist[endID2] {
		part1 = int(dist[endID1])
		dfs(endID1)
	} else {
		part1 = int(dist[endID2])
		dfs(endID2)
	}

	onBestPathPos := make(map[Pos]bool)
	for id := range onBestPath {
		// Extract position from ID
		pos := Pos{Y: id / (4 * Nx), X: (id / 4) % Nx}
		onBestPathPos[pos] = true
	}

	return part1, len(onBestPathPos)
}
