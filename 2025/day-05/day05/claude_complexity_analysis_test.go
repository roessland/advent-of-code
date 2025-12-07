package day05

import (
	"testing"
)

// TestPart2Complexity analyzes the actual behavior of Part2
func TestPart2Complexity(t *testing.T) {
	ranges, _ := ReadInput("input.txt")

	t.Logf("Initial ranges: %d", len(ranges))

	// Track the progression
	A := make([]Range, len(ranges))
	copy(A, ranges)
	B := []Range{}

	iteration := 0
	totalSubCalls := 0
	maxASize := len(A)

	for len(A) > 0 {
		iteration++
		b := A[len(A)-1]
		A = A[:len(A)-1]

		var nextA []Range
		subCalls := 0
		for _, a := range A {
			subCalls++
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

		B = append(B, b)
		A = nextA
		totalSubCalls += subCalls

		if len(A) > maxASize {
			maxASize = len(A)
		}

		// Log every 20 iterations
		if iteration%20 == 0 {
			t.Logf("Iteration %d: A size=%d, B size=%d, sub() calls=%d",
				iteration, len(A), len(B), subCalls)
		}
	}

	t.Logf("Final stats:")
	t.Logf("  Total iterations: %d", iteration)
	t.Logf("  Total Sub() calls: %d", totalSubCalls)
	t.Logf("  Max A size: %d", maxASize)
	t.Logf("  Final B size: %d", len(B))
	t.Logf("  Avg Sub() calls per iteration: %.2f", float64(totalSubCalls)/float64(iteration))
}

// TestPart1Complexity analyzes Part1 behavior
func TestPart1Complexity(t *testing.T) {
	ranges, ids := ReadInput("input.txt")

	t.Logf("Number of ranges: %d", len(ranges))
	t.Logf("Number of IDs to check: %d", len(ids))
	t.Logf("Total Contains() checks: %d", len(ranges)*len(ids))
	t.Logf("Complexity: O(n*m) = O(%d * %d) = O(%d)", len(ranges), len(ids), len(ranges)*len(ids))
}
