package main

import "fmt"

type Pair struct {
	Fst, Snd int32
}

func main() {
	input := []int32{10,16,6,0,1,17} // input
	//input := []int{0,3,6} // ex

	t := make([]Pair, 30000000)

	for turn := 1; turn < len(input)+1; turn++ {
		t[input[turn-1]] = Pair{Snd: int32(turn)}
	}

	var last = input[len(input)-1]
	for turn := len(input)+1; turn < 30000000+1; turn++ {
		if t[last].Fst == 0 {
			last = 0
		} else {
			last = t[last].Snd-t[last].Fst
		}
		t[last] = Pair{t[last].Snd, int32(turn)}
		if turn % 1000000 == 0 {
			fmt.Printf(".")
		}
	}
	fmt.Println("\n", last)
}