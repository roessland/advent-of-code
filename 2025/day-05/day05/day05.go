// Package day05 solves AoC 2025 Day 5
package day05

import (
	"embed"
	"fmt"

	"github.com/roessland/advent-of-code/2025/aocutil"
)

//go:embed input*.txt
var InputFS embed.FS

func ReadInput(inputName string) (ranges []Range, ids []int) {
	lines := aocutil.FSGetIntsInStringLines(InputFS, inputName)
	for _, line := range lines {
		if len(line) == 2 {
			ranges = append(ranges, Range{line[0], line[1]})
		} else if len(line) == 1 {
			ids = append(ids, line[0])
		}
	}
	return ranges, ids
}

type Range struct {
	A, B int
}

func (r Range) Len() int {
	return r.B - r.A + 1
}

func (r Range) Contains(id int) bool {
	return r.A <= id && id <= r.B
}

func (r Range) Sub(s Range) []Range {

	// Disjunct
	// ++r++    --s--
	// or
	// --s--    ++r++
	if r.B < s.A || s.B < r.A {
		return []Range{r}
	}
	// Contained in s
	//    ++++r++
	// -----s------
	if s.A <= r.A && r.B <= s.B {
		return nil
	}

	// Partial overlap 1
	//  ++++r++++
	//      ---s-----
	if r.A <= s.A && s.A <= r.B && r.B <= s.B {
		return []Range{{r.A, s.A - 1}}
	}

	// Partial overlap 2
	//    +++r++++
	// ---s-----
	if s.A <= r.A && r.A <= s.B && s.B <= r.B {
		return []Range{{s.B + 1, r.B}}
	}

	// Contains s
	// +++++r+++++
	//     --s--
	if r.A <= s.A && s.B <= r.B {
		return []Range{{r.A, s.A - 1}, {s.B + 1, r.B}}
	}

	fmt.Println("wat", r, s)
	panic("wat")
}

func IsFresh(ranges []Range, id int) bool {
	for _, ran := range ranges {
		if ran.Contains(id) {
			return true
		}
	}
	return false
}

func Part1(ranges []Range, ids []int) int {
	numFresh := 0
	for _, id := range ids {
		if IsFresh(ranges, id) {
			numFresh++
		}
	}
	return numFresh
}

func Part2(A []Range) int {
	// Take a range. Subtract it from all ranges in A, then move it to B.
	B := []Range{}
	for len(A) > 0 {
		// Pop range b from A
		b := A[len(A)-1]
		A = A[:len(A)-1]

		// Remove b from each range in A
		var nextA []Range
		for _, a := range A {
			aSubB := a.Sub(b)
			if len(aSubB) == 0 {
				continue
			} else if len(aSubB) == 1 {
				nextA = append(nextA, aSubB[0])
			} else if len(aSubB) == 2 {
				nextA = append(nextA, aSubB[0])
				nextA = append(nextA, aSubB[1])
			}
		}

		// Add b to B. Invariant: No overlap between elements in B.
		B = append(B, b)
		A = nextA
	}

	totalLen := 0
	for _, b := range B {
		totalLen += b.Len()
	}

	return totalLen
}
