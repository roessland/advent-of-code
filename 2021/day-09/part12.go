package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strings"
)

type Display struct {
	Inputs  []Signals
	Outputs []Signals
}

func main() {
	patterns := ReadInput()
	part1(patterns)
	part2(patterns)
}

func part1(displays []Display) {
	n := 0
	for _, display := range displays {
		for _, signals := range display.Outputs {
			if signals.Len() == 2 || signals.Len() == 4 || signals.Len() == 3 || signals.Len() == 7  { // 1
				n++
			}
		}
	}
	fmt.Println("Part 1:", n)
}

type Signals uint8

func SetOf(str string) Signals {
	var set Signals
	for _, c := range str {
		set = set | (1 << (c - 'a'))
	}
	return set
}

func (set Signals) Len() int {
	return bits.OnesCount8(uint8(set))
}

func Decode(p map[int]Signals, signalsList []Signals) int {
	digitForSignals := make(map[Signals]int)
	for digit, signals := range p {
		digitForSignals[signals] = digit
	}

	out := 0
	for _, signals := range signalsList {
		out = out*10 + digitForSignals[signals]
	}
	return out
}

func part2(displays []Display) {
	totalOutput := 0
	for _, display := range displays {
		signalsFor := make(map[int]Signals)

		// First pass -- Get 1, 4, 7 since they are uniquely identified by
		// the number of segments only.
		for _, signals := range display.Inputs {
			if signals.Len() == 2 { // 1
				signalsFor[1] = signals
			} else if signals.Len() == 4 { // 4
				signalsFor[4] = signals
			} else if signals.Len() == 3 { // 7
				signalsFor[7] = signals
			} else if signals.Len() == 7 { // 8
				signalsFor[8] = signals
			}
		}

		// Second pass -- Moderately clever logic based on set operations.
		for _, signals := range display.Inputs {
			switch signals {
			case signalsFor[1], signalsFor[4], signalsFor[7], signalsFor[8]:
				continue
			}
			Len := signals.Len()
			and7 := (signalsFor[7] & signals).Len()
			or7 := (signalsFor[7] | signals).Len()
			and4 := (signalsFor[4] & signals).Len()

			if Len == 6 && and7 == 3 && or7 == 6 && and4 == 3 {
				signalsFor[0] = signals
			} else if Len == 5 && and7 == 2 && or7 == 6 && and4 == 2 {
				signalsFor[2] = signals
			} else if Len == 5 && and7 == 3  {
				signalsFor[3] = signals
			} else if Len == 5 && and7 == 2 && or7 == 6 && and4 == 3  {
				signalsFor[5] = signals
			} else if Len == 6 && and7 == 2 {
				signalsFor[6] = signals
			} else if Len == 6 && and7 == 3 && or7 == 6 && and4 == 4 {
				signalsFor[9] = signals
			}
		}

		// Decode output
		totalOutput += Decode(signalsFor, display.Outputs)
	}
	fmt.Println("Part 2:", totalOutput)
}

func ReadInput() []Display {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var patterns []Display
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		pat := Display{}
		for _, signalsStr := range strings.Split(parts[0], " ") {
			pat.Inputs = append(pat.Inputs, SetOf(signalsStr))
		}
		for _, signalsStr := range strings.Split(parts[1], " ") {
			pat.Outputs = append(pat.Outputs, SetOf(signalsStr))
		}
		patterns = append(patterns, pat)
	}
	return patterns
}
