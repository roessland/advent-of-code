package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var Pad = [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
}

const numRows int = 3
const numCols int = 3

type State struct {
	m int
	n int
}

func (s State) Move(direction rune) State {
	switch direction {
	case 'R':
		return State{s.m, Clamp(s.n+1, 0, numCols-1)}
	case 'U':
		return State{Clamp(s.m-1, 0, numRows-1), s.n}
	case 'L':
		return State{s.m, Clamp(s.n-1, 0, numCols-1)}
	case 'D':
		return State{Clamp(s.m+1, 0, numRows-1), s.n}
	default:
		panic("No such direction")
	}
}

func (s State) GetNumber() int {
	return Pad[s.m][s.n]
}

func main() {
	s := State{1, 1}
	code := make([]int, 0)
	buf, _ := ioutil.ReadFile("input1.txt")
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		for _, direction := range line {
			s = s.Move(direction)
		}
		code = append(code, s.GetNumber())
	}
	fmt.Println("The code is: ", code)
}

func Clamp(n, lo, hi int) int {
	if n <= lo {
		return lo
	} else if n >= hi {
		return hi
	} else {
		return n
	}
}
