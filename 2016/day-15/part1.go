package main

import (
	"fmt"
)

type Disc struct {
	Num          int
	NumPositions int
	Offset       int
}

func main() {
	discs := []Disc{
		Disc{1, 17, 15},
		Disc{2, 3, 2},
		Disc{3, 19, 4},
		Disc{4, 13, 2},
		Disc{5, 7, 2},
		Disc{6, 5, 0},
		//Disc{7, 11, 0}, // uncomment for part 1
	}

	for t := 1; ; t += 17 {
		canPress := true
		for _, disc := range discs {
			if (t+disc.Num+disc.Offset)%disc.NumPositions != 0 {
				canPress = false
				break
			}
		}
		if canPress {
			fmt.Println(t)
			break
		}
	}
}
