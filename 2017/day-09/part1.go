package main

import "fmt"
import "log"
import "io"
import "os"
import "bufio"

const (
	groupState = iota
	garbageState
	closedState
	groupStateExpectOpen
	groupStateNormal
	garbageStateNormal
	garbageStateIgnoreOne
)

type StreamParser struct {
	depth          int
	state          int
	groupState     int
	garbageState   int
	rr             io.RuneReader
	totalScore     int
	garbageRemoved int
}

func NewStreamParser(rr io.RuneReader) *StreamParser {
	sp := StreamParser{0, groupState, groupStateExpectOpen, garbageStateNormal, rr, 0, 0}
	return &sp
}

func (sp *StreamParser) parseGroup(r rune) {
	switch sp.groupState {
	case groupStateExpectOpen:
		if r != '{' {
			log.Fatalf("Expected { but got %c", r)
		}
		sp.depth++
		sp.groupState = groupStateNormal
	case groupStateNormal:
		if r == '{' {
			sp.depth++
			sp.groupState = groupStateNormal
		} else if r == '<' {
			sp.state = garbageState
			sp.groupState = groupStateNormal
		} else if r == '}' {
			sp.totalScore += sp.depth
			sp.depth--
			sp.groupState = groupStateNormal
		}
	}

}

func (sp *StreamParser) parseGarbage(r rune) {
	switch sp.garbageState {
	case garbageStateNormal:
		switch r {
		case '>':
			sp.state = groupState
			sp.groupState = groupStateNormal
			sp.garbageState = garbageStateNormal
		case '!':
			sp.garbageState = garbageStateIgnoreOne
		default:
			sp.garbageRemoved++
		}
	case garbageStateIgnoreOne:
		sp.garbageState = garbageStateNormal
	default:
		log.Fatalf("Unknown garbage state %d", sp.garbageState)
	}
}

func (sp *StreamParser) FindTotalScore() int {
	for {
		r, _, err := sp.rr.ReadRune()
		if err == io.EOF {
			if sp.depth != 0 {
				log.Fatalf("Unexpected EOF; depth was %d", sp.depth)
			}
			return sp.totalScore
		}
		if err != nil {
			log.Fatal(err)
		}
		switch sp.state {
		case groupState:
			sp.parseGroup(r)
		case garbageState:
			sp.parseGarbage(r)
		case closedState:
			log.Fatal("Stream is already closed.")
		default:
			log.Fatalf("Unexpected state encountered: %v", sp.state)
		}
	}
}

func main() {
	sp := NewStreamParser(bufio.NewReader(os.Stdin))
	totalScore := sp.FindTotalScore()
	fmt.Println("Total score for all groups:", totalScore)
	fmt.Println("Garbage removed:", sp.garbageRemoved)
}
