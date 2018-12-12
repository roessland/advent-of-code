package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const Nmax int = 1000

type State struct {
	Generation int
	Offset     int
	Data       []byte
}

func (s0 State) Next(rules map[string]byte) State {
	s1 := State{
		Generation: s0.Generation + 1,
		Offset:     s0.Offset,
		Data:       make([]byte, len(s0.Data)),
	}
	for i := range s1.Data {
		s1.Data[i] = '.'
	}
	for i := 2; i < len(s1.Data)-2; i++ {
		s1.Data[i] = Apply(rules, string(s0.Data[i-2:i+3]))
	}
	return s1
}

func Apply(rules map[string]byte, key string) byte {
	if val, ok := rules[key]; ok {
		return val
	} else {
		return '.'
	}
}

func ReadRules(lines []string) map[string]byte {
	rules := map[string]byte{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " => ")
		rules[parts[0]] = parts[1][0]
	}
	return rules
}

func Count(state State) int {
	sum := 0
	for i, val := range state.Data {
		if val == '#' {
			sum += i - Nmax
		}
	}
	return sum
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\n")
	initial := []byte(strings.Trim(strings.Split(lines[0], ": ")[1], " \n"))
	rules := ReadRules(lines[2:])

	state := State{Generation: 0, Offset: Nmax, Data: make([]byte, 2*Nmax+len(initial))}
	for i := range state.Data {
		state.Data[i] = '.'
	}
	copy(state.Data[Nmax:Nmax+len(initial)], initial)
	prevCount := 0
	diff := 0
	aNiceNumber := 500
	for state.Generation < aNiceNumber {
		state = state.Next(rules)
		count := Count(state)
		diff = count - prevCount
		prevCount = count
	}

	fmt.Println((50000000000-aNiceNumber)*diff + prevCount)
}
