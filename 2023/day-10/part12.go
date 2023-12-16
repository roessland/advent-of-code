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

func (g Grid) Neighbors(p Pos) []Pos {
	neighbors := []Pos{}
	n := p.North()
	if strings.IndexByte("|7FS", g.At(n)) != -1 {
		neighbors = append(neighbors, n)
	}

	s := p.South()
	if strings.IndexByte("|LJS", g.At(s)) != -1 {
		neighbors = append(neighbors, s)
	}

	w := p.West()
	if strings.IndexByte("-LFS", g.At(w)) != -1 {
		neighbors = append(neighbors, w)
	}

	e := p.East()
	if strings.IndexByte("-7JS", g.At(e)) != -1 {
		neighbors = append(neighbors, e)
	}

	return neighbors
}

func part1(g Grid) {
	posS := g.FindStartPos()
	ans := MaxLoopDist(g, posS)
	fmt.Println("max loop dist:", ans)
}

func MaxLoopDist(g Grid, startPos Pos) int {
	dists := map[Pos]int{startPos: 0}
	todo := map[Pos]struct{}{startPos: {}}
	var bfs func()
	bfs = func() {
		// Pop tile closest to start
		p := argMin(keys(todo), func(p Pos) int { return dists[p] })
		delete(todo, p)

		// Add all unvisited neighbors to todo
		neighbors := g.Neighbors(p)
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
	return maxVal(dists)
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
	// how do I functional, without a looo?
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
