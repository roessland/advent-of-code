package main

import (
	"bufio"
	"fmt"
	"os"
)

type State [][]byte

func Copy(state State) State {
	c := make(State, len(state))
	for i := range state {
		c[i] = make([]byte, len(state[0]))
		copy(c[i], state[i])
	}
	return c
}

func Inc(s [][]int, i, j int) {
	if i < 0  || i >= len(s) {
		return
	}
	if j < 0 || j >= len(s[0]) {
		return
	}
	s[i][j]++
}

func (s State) Count() [][]int {
	cs := make([][]int, len(s))
	for i := range s {
		cs[i] = make([]int, len(s[0]))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			if s[i][j] != '#' {
				continue
			}
			Inc(cs, i-1, j)
			Inc(cs, i+1, j)
			Inc(cs, i, j+1)
			Inc(cs, i, j-1)
		}
	}
	return cs
}

func (s State) Evolve() {
	cs := s.Count()
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			if s[i][j] == '#' {
				if cs[i][j] == 1 {
					continue
				}
				s[i][j] = '.'
			} else {
				if cs[i][j] == 1 || cs[i][j] == 2 {
					s[i][j] = '#'
				}
			}
		}
	}
}

func (s State) Print() {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			fmt.Printf("%c", s[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s State) Int() (val int) {
	k := 1
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			if s[i][j] == '#' {
				val += k
			}
			k *= 2
		}
	}
	return val
}

func main() {
	var state State
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		state = append(state, []byte(line))
	}

	freq := map[int]int{}
	for i := 0; i < 5000; i++ {
		rating := state.Int()
		freq[rating]++
		if freq[rating] == 2 {
			fmt.Println("Part 1:", rating)
			break
		}
		state.Evolve()
	}
}