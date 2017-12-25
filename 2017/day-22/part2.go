package main

import "fmt"
import "os"
import "bufio"

const (
	Clean int = iota
	Weakened
	Infected
	Flagged
)

func NextState(s int) int {
	return (s + 1) % 4
}

type Vec struct {
	X, Y int
}

func (p Vec) Add(v Vec) Vec {
	return Vec{p.X + v.X, p.Y + v.Y}
}

func (v Vec) Neg() Vec {
	return Vec{-v.X, -v.Y}
}

// CCW=1 is counterclockwise, CCW=-1 is clockwise
func Rotate(v Vec, CCW int) Vec {
	if v.Y == -1 {
		return Vec{-CCW, 0}
	}
	if v.Y == 1 {
		return Vec{CCW, 0}
	}
	if v.X == -1 {
		return Vec{0, CCW}
	}
	if v.X == 1 {
		return Vec{0, -CCW}
	}
	return Vec{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	state := map[Vec]int{}
	i, j := 0, 0
	for scanner.Scan() {
		row := scanner.Text()
		i = 0
		for _, c := range row {
			if c == '#' {
				state[Vec{i, j}] = Infected
			}
			i++
		}
		j++
	}
	pos := Vec{i / 2, j / 2}
	dir := Vec{0, -1}
	infectionsCaused := 0
	for t := 0; t < 10000000; t++ {
		switch state[pos] {
		case Clean:
			dir = Rotate(dir, 1)
		case Weakened:
		case Infected:
			dir = Rotate(dir, -1)
		case Flagged:
			dir = dir.Neg()
		}
		state[pos] = NextState(state[pos])
		if state[pos] == Infected {
			infectionsCaused++
		}
		pos = pos.Add(dir)
	}
	fmt.Println(infectionsCaused)
}
