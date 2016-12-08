package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ArgMin(freq map[rune]int) rune {
	if len(freq) == 0 {
		panic("No data")
	}
	minFreq := 9999999999
	minRune := '_'
	for c, f := range freq {
		if f < minFreq {
			minFreq = f
			minRune = c
		}
	}
	return minRune
}

func ArgMax(freq map[rune]int) rune {
	if len(freq) == 0 {
		panic("No data")
	}
	maxFreq := 0
	maxRune := '_'
	for c, f := range freq {
		if f > maxFreq {
			maxFreq = f
			maxRune = c
		}
	}
	return maxRune
}

type Message struct {
	frequencies []map[rune]int
}

func NewMessage(length int) Message {
	m := Message{make([]map[rune]int, length)}
	for i, _ := range m.frequencies {
		m.frequencies[i] = make(map[rune]int)
	}
	return m
}

func (m Message) Add(msg string) {
	for i, c := range msg {
		m.frequencies[i][c]++
	}
}

func (m Message) ErrorCorrected(typ string) string {
	msg := make([]rune, len(m.frequencies))
	for i, freq := range m.frequencies {
		if typ == "max" {
			msg[i] = ArgMax(freq)
		} else if typ == "min" {
			msg[i] = ArgMin(freq)
		} else {
			panic("oh snap")
		}
	}
	return string(msg)
}

func main() {
	m := NewMessage(8)
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		m.Add(line)
	}
	fmt.Println("Max freq:", m.ErrorCorrected("max"))
	fmt.Println("Min freq:", m.ErrorCorrected("min"))
}
