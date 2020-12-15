package main

import "fmt"

func main() {
	input := []int{10,16,6,0,1,17} // input
	//input := []int{0,3,6} // ex

	t0 := map[int]int{}
	t1 := map[int]int{}

	for turn := 1; turn < len(input)+1; turn++ {
		t1[input[turn-1]] = turn
	}

	last := input[len(input)-1]
	for turn := len(input)+1; turn < 30000000+1; turn++ {
		if t0[last] == 0 {
			last = 0
		} else {
			last = t1[last] - t0[last]
		}
		t0[last], t1[last] = t1[last], turn
		if turn % 1000000 == 0 {
			fmt.Printf(".")
		}
	}
	fmt.Println("\n", last)
}