package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type State [100]byte

type Pos struct {
	I, J int
}

func (state State) At(pos Pos) byte {
	return state[pos.J*10+pos.I]
}

func (state *State) Inc(pos Pos) {
	state[pos.J*10+pos.I]++
}

func (state *State) Zero(pos Pos) {
	state[pos.J*10+pos.I] = 0
}

func (state State) AllPositions() <-chan Pos {
	ch := make(chan Pos)
	go func() {
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				ch <- Pos{i, j}
			}
		}
		close(ch)
	}()
	return ch
}

func (state State) AdjacentPositions(pos Pos) []Pos {
	var legalAdjPositions []Pos
	for _, adjPos := range []Pos{
		{pos.I - 1, pos.J}, {pos.I + 1, pos.J},
		{pos.I, pos.J - 1}, {pos.I, pos.J + 1},
		{pos.I - 1, pos.J - 1}, {pos.I + 1, pos.J - 1},
		{pos.I - 1, pos.J + 1}, {pos.I + 1, pos.J + 1},
	} {
		if adjPos.I < 0 || adjPos.I >= 10 {
			continue
		}
		if adjPos.J < 0 || adjPos.J >= 10 {
			continue
		}
		legalAdjPositions = append(legalAdjPositions, adjPos)
	}
	return legalAdjPositions
}

func main() {
	m := ReadInput()
	part12(m)
}

func part12(state State) {
	var totalFlashes int
	for i := 1; ; i++ {
		var flashes int
		state, flashes = update(state)

		// Part 1
		totalFlashes += flashes
		if i == 100 {
			fmt.Println(totalFlashes)
		}

		// Part 2
		if flashes == 100 {
			fmt.Println(i)
			break
		}
	}

}

func update(state State) (State, int) {
	for pos := range state.AllPositions() {
		state.Inc(pos)
	}

	flashed := State{}
	for {
		done := true
		for pos := range state.AllPositions() {
			level := state.At(pos)
			if level > 9 && flashed.At(pos) == 0 {
				done = false
				state.Zero(pos)
				flashed.Inc(pos)
				for _, adjPosition := range state.AdjacentPositions(pos) {
					state.Inc(adjPosition)
				}
			}
		}
		if done {
			break
		}
	}

	flashes := 0
	for pos := range state.AllPositions() {
		if flashed.At(pos) != 0 {
			flashes++
			state.Zero(pos)
		}
	}
	return state, flashes
}

func ReadInput() State {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	state := State{}
	for j, line := range lines {
		for i, c := range line {
			state[j*10+i] = byte(c - '0')
		}
	}
	return state
}
