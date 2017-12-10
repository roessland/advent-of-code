package main

// Please excuse the ugly code, <insert excuses here>

import "fmt"
import "os"

// 1 3 5 7 9 11 13 15 17 19 21 23 ...

// 37 36 35 34 33 32 31
// 38 17 16 15 14 13 30
// 39 18  5  4  3 12 29
// 40 19  6  1  2 11 28
// 41 20  7  8  9 10 27
// 42 21 22 23 24 25 26
// 43 44 45 46 47 48 49

const (
	West = iota
	North
	East
	South
)

func turn(dx0, dy0 int) (int, int) {
	if dx0 == 1 {
		return 0, -1
	} else if dx0 == -1 {
		return 0, 1
	} else if dy0 == 1 {
		return 1, 0
	} else if dy0 == -1 {
		return -1, 0
	}
	return -999, -999
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(n int) {
	if n > 347991 {
		fmt.Println("First larger value is", n)
		os.Exit(0)
	}
}

func NeighborSum(m map[Pos]int, p Pos) int {
	sum := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			sum += m[Pos{p[0] + dx, p[1] + dy}]
		}
	}
	return sum
}

type Pos [2]int

func (p Pos) E() Pos {
	return Pos{p[0] + 1, p[1]}
}
func (p Pos) N() Pos {
	return Pos{p[0], p[1] - 1}
}
func (p Pos) W() Pos {
	return Pos{p[0] - 1, p[1]}
}
func (p Pos) S() Pos {
	return Pos{p[0], p[1] + 1}
}

func main() {
	squares := map[Pos]int{}
	squares[Pos{0, 0}] = 1
	x, y := 0, 0
	dx, dy := 1, 0
	_, _, _, _ = x, y, dx, dy
	r := -1
	for {
		r += 2
		fmt.Println("Going right for", r)
		for i := 0; i < r; i++ {
			x += dx
			y += dy
			n := NeighborSum(squares, Pos{x, y})
			squares[Pos{x, y}] = n
			check(n)
		}
		dx, dy = turn(dx, dy)
		fmt.Println("Going up for", r)
		for i := 0; i < r; i++ {
			x += dx
			y += dy
			n := NeighborSum(squares, Pos{x, y})
			squares[Pos{x, y}] = n
			check(n)
		}
		dx, dy = turn(dx, dy)
		fmt.Println("Going left for", r+1)
		for i := 0; i < r+1; i++ {
			x += dx
			y += dy
			n := NeighborSum(squares, Pos{x, y})
			squares[Pos{x, y}] = n
			check(n)
		}
		dx, dy = turn(dx, dy)
		fmt.Println("Going down for", r+1)
		for i := 0; i < r+1; i++ {
			x += dx
			y += dy
			n := NeighborSum(squares, Pos{x, y})
			squares[Pos{x, y}] = n
			check(n)
		}
		dx, dy = turn(dx, dy)
	}
}
