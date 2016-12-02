package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var Pad = [][]rune{
	[]rune{' ', ' ', '1', ' ', ' '},
	[]rune{' ', '2', '3', '4', ' '},
	[]rune{'5', '6', '7', '8', '9'},
	[]rune{' ', 'A', 'B', 'C', ' '},
	[]rune{' ', ' ', 'D', ' ', ' '},
}

const numRows int = 5
const numCols int = 5

type State struct {
	m int
	n int
}

func (s State) Move(direction rune) State {
	var next State
	switch direction {
	case 'R':
		next = State{s.m, Clamp(s.n+1, 0, numCols-1)}
	case 'U':
		next = State{Clamp(s.m-1, 0, numRows-1), s.n}
	case 'L':
		next = State{s.m, Clamp(s.n-1, 0, numCols-1)}
	case 'D':
		next = State{Clamp(s.m+1, 0, numRows-1), s.n}
	default:
		panic("No such direction")
	}
	if next.GetRune() == ' ' {
		return s
	} else {
		return next
	}
}

func (s State) GetRune() rune {
	return Pad[s.m][s.n]
}

func main() {
	s := State{2, 0}
	code := make([]rune, 0)
	buf, _ := ioutil.ReadFile("input1.txt")
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		for _, direction := range line {
			s = s.Move(direction)
		}
		code = append(code, s.GetRune())
	}
	fmt.Println("The code is: ", string(code))
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
