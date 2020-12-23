package main

import (
	"fmt"
)

const N = 1000000


type Cup struct {
	Label int
	Prev *Cup
	Next *Cup
}

func NewRing(labels ...int) ([]Cup, *Cup) {
	cups := make([]Cup, N)
	for i := 0; i < N; i++ {
		var label, labelNext int
		if i == 0 {
			label = labels[i]
			labelNext = labels[i+1]
		} else if i < len(labels)-1 {
			label = labels[i]
			labelNext = labels[i+1]
		} else if i == len(labels)-1 {
			label = labels[i]
			labelNext = len(labels) + 1
		} else if i == len(labels) {
			label = i+1
			labelNext = i +2
		} else if i < N-1 {
			label = i+1
			labelNext = i+2
		} else {
			label = i+1
			labelNext = labels[0]
		}
		cups[label-1] = Cup{
			Label: label,
			Next: &cups[labelNext-1],
		}
	}
	return cups, &cups[labels[0]-1]
}

func main() {
	cups, curr := NewRing(4,6,9,2,1,7,5,3,8) // input
	//cups, curr := NewRing(3,8,9,1,2,5,4,6,7) // example

	for i := 0; i < 10*N; i++ {
		// Take three cups out from the ring
		cup1 := curr.Next
		cup2 := cup1.Next
		cup3 := cup2.Next
		curr.Next = cup3.Next
		cup3.Next = nil

		// Choose a destination label
		dstLabel := curr.Label-1
		if dstLabel == 0 {
			dstLabel = N
		}
		for dstLabel == cup1.Label || dstLabel == cup2.Label || dstLabel == cup3.Label {
			dstLabel--
			if dstLabel == 0 {
				dstLabel = N
			}
		}

		// Locate destination cup
		dst := &cups[dstLabel-1]

		// Insert the three cups clockwise of destination cup
		cup3.Next = dst.Next
		dst.Next = cup1

		// Chose next current cup as the one
		curr = curr.Next
	}

	one := cups[1-1]

	fmt.Println(one.Next.Label * one.Next.Next.Label)
}

