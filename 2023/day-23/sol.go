package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

type Map []string

func (m Map) Width() int {
	return len(m[0])
}

func (m Map) Height() int {
	return len(m)
}

func (m Map) At(p Pos) byte {
	y, x := p.Y, p.X
	if y < 0 || y >= m.Height() || x < 0 || x >= m.Width() {
		return '#'
	}
	return m[y][x]
}

type Pos struct {
	Y, X int
}

func (p Pos) Up() Pos {
	return Pos{p.Y - 1, p.X}
}

func (p Pos) Down() Pos {
	return Pos{p.Y + 1, p.X}
}

func (p Pos) Left() Pos {
	return Pos{p.Y, p.X - 1}
}

func (p Pos) Right() Pos {
	return Pos{p.Y, p.X + 1}
}

func (p Pos) Neighbors(currTile byte) []Pos {
	switch currTile {
	case '>':
		return []Pos{p.Right()}
	case '<':
		return []Pos{p.Left()}
	case '^':
		return []Pos{p.Up()}
	case 'v':
		return []Pos{p.Down()}
	default:
		return []Pos{p.Up(), p.Down(), p.Left(), p.Right()}
	}
}

type Graph map[Pos]Edges

type Edges map[Pos]int

func MakeGraph(m Map) Graph {
	g := make(Graph)
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			p := Pos{y, x}
			if m.At(p) == '#' {
				continue
			}
			g[Pos{y, x}] = make(Edges)
		}
	}

	for p := range g {
		for _, n := range p.Neighbors2(m.At(p)) {
			if m.At(n) == '#' {
				continue
			}
			g[p][n] = 1
		}
	}

	return g
}

type Edge struct {
	To   Pos
	Dist int
}

func (g Graph) Compress() {
	for p, edges := range g {
		if len(edges) != 2 {
			continue
		}
		es := make([]Edge, 0)
		for n, dist := range edges {
			es = append(es, Edge{n, dist})
		}
		// ┌─┐◀─┌─┐◀─┌─┐
		// │a│─▶│p│─▶│b│
		// └─┘  └─┘  └─┘
		pToA, pToB := es[0], es[1]
		a, b := pToA.To, pToB.To

		// Link a to b and b to a
		g[a][b] = pToA.Dist + pToB.Dist
		g[b][a] = pToA.Dist + pToB.Dist

		// Remove p and edges to it from graph
		delete(g[a], p)
		delete(g[b], p)
		delete(g, p)
	}
}

func (p Pos) Neighbors2(currTile byte) []Pos {
	return []Pos{p.Up(), p.Down(), p.Left(), p.Right()}
}

func LongestLen(g Graph, curr Pos, target Pos, visited map[Pos]bool) int {
	if curr == target {
		return 0
	}

	visited[curr] = true
	defer delete(visited, curr)

	longest := -1
	for neighborPos, distToNeighbor := range g[curr] {
		if visited[neighborPos] {
			continue
		}
		viaNeigh := LongestLen(g, neighborPos, target, visited)
		if viaNeigh != -1 && distToNeighbor+viaNeigh > longest {
			longest = distToNeighbor + viaNeigh
		}
	}
	return longest
}

var visited map[Pos]bool

func main() {
	m := ReadInput()
	pA := Pos{Y: 0, X: strings.Index(m[0], ".")}
	pB := Pos{Y: m.Height() - 1, X: strings.Index(m[m.Height()-1], ".")}
	if pA.X == -1 || pB.X == -1 {
		panic("No start or end found")
	}

	g := MakeGraph(m)

	for {
		before := len(g)
		g.Compress()
		after := len(g)
		fmt.Println("Compressing graph:", before, after)
		if before == after {
			break
		}
	}
	fmt.Println(LongestLen(g, pA, pB, map[Pos]bool{}))

	// // p := tea.NewProgram(initialModel(bricks))
	// p := tea.NewProgram(initialModel(&m, visited), tea.WithAltScreen())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Error running program: %s\n", err)
	// 	os.Exit(1)
	// }
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ReadInput() Map {
	// m := aocutil.ReadLines("input-ex1.txt")
	m := aocutil.ReadLines("input.txt")
	return m
}
