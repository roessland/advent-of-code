package main1

import (
	"fmt"
)

type Cup struct {
	Label int
	Prev *Cup
	Next *Cup
}

func NewRing(labels ...int) *Cup {
	cups := make([]Cup, len(labels))
	for i, label := range labels {
		cups[i] = Cup{
			Label: label,
			Prev: &cups[(i-1+len(cups)) % len(cups)],
			Next: &cups[(i+1)%len(cups)],
		}
	}
	return &cups[0]
}

func main() {
	curr := NewRing(4,6,9,2,1,7,5,3,8) // input
	//curr := NewRing(3,8,9,1,2,5,4,6,7) // example

	for i := 0; i < 100; i++ {
		// Take three cups out from the ring
		cup1 := curr.Next
		cup2 := cup1.Next
		cup3 := cup2.Next
		curr.Next = cup3.Next
		cup3.Next = nil

		// Choose a destination label
		dstLabel := (curr.Label-1-1+9) % 9 + 1
		for dstLabel == cup1.Label || dstLabel == cup2.Label || dstLabel == cup3.Label {
			dstLabel = (dstLabel-1-1+9) % 9 + 1
		}

		// Locate destination cup
		dst := curr
		for dst.Label != dstLabel {
			dst = dst.Next
		}

		// Insert the three cups clockwise of destination cup
		cup3.Next = dst.Next
		dst.Next = cup1

		// Chose next current cup as the one
		curr = curr.Next
	}

	one := curr
	for one.Label != 1 {
		one = one.Next
	}

	curr = one.Next
	for curr != one {
		fmt.Print(curr.Label)
		curr = curr.Next
	}
	fmt.Println()
}

