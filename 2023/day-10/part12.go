package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

type Grid []string

func ReadInput() Grid {
	return aocutil.ReadLines("input.txt")
}

func main() {
	input := ReadInput()
	part1(input)
	part2(input)
}

type Pos struct {
	X, Y int
}

func (p Pos) North() Pos {
	return Pos{p.X, p.Y - 1}
}

func (p Pos) South() Pos {
	return Pos{p.X, p.Y + 1}
}

func (p Pos) East() Pos {
	return Pos{p.X + 1, p.Y}
}

func (p Pos) West() Pos {
	return Pos{p.X - 1, p.Y}
}

func (g Grid) At(p Pos) byte {
	x, y := p.X, p.Y
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return '.'
	}
	return g[y][x]
}

func (g Grid) FindStartPos() Pos {
	x := strings.IndexRune(g[0], 'S')
	if x == -1 {
		startPos := g[1:].FindStartPos()
		return Pos{startPos.X, startPos.Y + 1}
	}
	return Pos{x, 0}
}

func (g Grid) ConnectedNeighbors(p Pos) []Pos {
	neighbors := []Pos{}
	n := p.North()
	if strings.IndexByte("|7FS", g.At(n)) != -1 &&
		strings.IndexByte("|LJS", g.At(p)) != -1 {
		neighbors = append(neighbors, n)
	}

	s := p.South()
	if strings.IndexByte("|LJS", g.At(s)) != -1 &&
		strings.IndexByte("|F7S", g.At(p)) != -1 {
		neighbors = append(neighbors, s)
	}

	w := p.West()
	if strings.IndexByte("-LFS", g.At(w)) != -1 &&
		strings.IndexByte("-7JS", g.At(p)) != -1 {
		neighbors = append(neighbors, w)
	}

	e := p.East()
	if strings.IndexByte("-7JS", g.At(e)) != -1 &&
		strings.IndexByte("-FLS", g.At(p)) != -1 {
		neighbors = append(neighbors, e)
	}

	return neighbors
}

func part1(g Grid) {
	posS := g.FindStartPos()
	ans, _ := MaxLoopDist(g, posS)
	fmt.Println("max loop dist:", ans)
}

func MaxLoopDist(g Grid, startPos Pos) (int, map[Pos]int) {
	dists := map[Pos]int{startPos: 0}
	todo := map[Pos]struct{}{startPos: {}}
	var bfs func()
	bfs = func() {
		// Pop tile closest to start
		p := argMin(keys(todo), func(p Pos) int { return dists[p] })
		delete(todo, p)

		// Add all unvisited neighbors to todo
		neighbors := g.ConnectedNeighbors(p)
		forEach(neighbors, func(n Pos) {
			if _, ok := dists[n]; !ok {
				dists[n] = dists[p] + 1
				todo[n] = struct{}{}
			}
		})

		// Repeat until all tiles have been visited
		if len(todo) > 0 {
			bfs()
		}
	}
	bfs()
	return maxVal(dists), dists
}

func part2(g Grid) {
	p0 := g.FindStartPos()
	p1 := g.ConnectedNeighbors(p0)[0]
	ingress := map[Pos]byte{}
	egress := map[Pos]byte{}
	var dfs func(p Pos, prev Pos)
	dfs = func(p, prev Pos) {
		// If already visited, done
		if _, ok := ingress[p]; ok {
			return
		}

		// Visit this one
		dir := Pos{p.X - prev.X, p.Y - prev.Y}
		switch dir {
		case Pos{0, -1}:
			ingress[p] = 'v'
			egress[prev] = '^'
		case Pos{0, 1}:
			ingress[p] = '^'
			egress[prev] = 'v'
		case Pos{-1, 0}:
			ingress[p] = '>'
			egress[prev] = '<'
		case Pos{1, 0}:
			ingress[p] = '<'
			egress[prev] = '>'
		}

		// Find next
		var next Pos
		neighbors := g.ConnectedNeighbors(p)
		if neighbors[0] == prev {
			next = neighbors[1]
		} else {
			next = neighbors[0]
		}

		// Visit next
		dfs(next, p)
	}
	dfs(p1, p0)

	isRightOf := map[Pos]bool{}
	for {
		initialLen := len(isRightOf)
		for y := 0; y < len(g); y++ {
			for x := 0; x < len(g[y]); x++ {
				p := Pos{x, y}
				// Skip loop tiles
				if _, isLoop := egress[p]; isLoop {
					continue
				}

				// ingress      egress
				//   <     J     <
				// v o ^  |.L  v   ^
				//   >     -     >

				//
				// o ^ .
				//
				e := p.East()
				if egress[e] == '^' || ingress[e] == 'v' {
					isRightOf[p] = true
				}

				// . v o
				w := p.West()
				if egress[w] == 'v' || ingress[w] == '^' {
					isRightOf[p] = true
				}

				// .
				// <
				// o
				n := p.North()
				if egress[n] == '<' || ingress[n] == '>' {
					isRightOf[p] = true
				}

				// o
				// <
				// .
				s := p.South()
				if egress[s] == '>' || ingress[s] == '<' {
					isRightOf[p] = true
				}

				// Flood fill. If a neighbor is on the left, this tile is on the left.
				if isRightOf[e] || isRightOf[w] || isRightOf[n] || isRightOf[s] {
					isRightOf[p] = true
				}
			}
		}
		// Flood fill is done
		if len(isRightOf) == initialLen {
			break
		}
	}

	fmt.Println("\n\nIngress:")
	for y := 0; y < len(g); y++ {
		fmt.Println()
		for x := 0; x < len(g[y]); x++ {
			p := Pos{x, y}
			if _, ok := ingress[p]; ok {
				fmt.Printf("%c", ingress[p])
				// fmt.Printf("%c", g.At(p))
			} else if isRightOf[p] {
				fmt.Printf("O")
			} else {
				fmt.Printf("%c", g.At(p))
			}
		}
	}

	fmt.Println("\n", len(isRightOf))
}

func argMin[T comparable](xs []T, f func(T) int) T {
	arg, _ := argMinHelper(xs, f)
	return arg
}

func argMinHelper[T comparable](xs []T, f func(T) int) (T, int) {
	if len(xs) == 0 {
		panic("argMin of empty list")
	}
	if len(xs) == 1 {
		return xs[0], f(xs[0])
	}
	argHead, rest := xs[0], xs[1:]
	valHead := f(argHead)
	argMinRest, minRest := argMinHelper(rest, f)
	if valHead < minRest {
		return argHead, valHead
	}
	return argMinRest, minRest
}

func forEach(xs []Pos, f func(Pos)) {
	if len(xs) == 0 {
		return
	}
	f(xs[0])
	forEach(xs[1:], f)
}

func maxVal[T comparable](m map[T]int) int {
	return max(vals(m))
}

func vals[T comparable](m map[T]int) []int {
	vs := []int{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func keys[T comparable, V any](m map[T]V) []T {
	ks := []T{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func max(xs []int) int {
	if len(xs) == 0 {
		panic("max of empty list")
	}
	if len(xs) == 1 {
		return xs[0]
	}
	head, rest := xs[0], xs[1:]
	maxRest := max(rest)
	if head > maxRest {
		return head
	}
	return maxRest
}
