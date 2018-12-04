package main

import "fmt"

func main() {
	set := make(map[int]int)
	slot := 0
	state := "A"
	for i := 0; i < 12481997; i++ {
		switch state {
		case "A":
			if set[slot] == 0 {
				set[slot] = 1
				slot++
				state = "B"
			} else {
				set[slot] = 0
				slot--
				state = "C"
			}
		case "B":
			if set[slot] == 0 {
				set[slot] = 1
				slot--
				state = "A"
			} else {
				set[slot] = 1
				slot++
				state = "D"
			}
		case "C":
			if set[slot] == 0 {
				set[slot] = 0
				slot--
				state = "B"
			} else {
				set[slot] = 0
				slot--
				state = "E"
			}
		case "D":
			if set[slot] == 0 {
				set[slot] = 1
				slot++
				state = "A"
			} else {
				set[slot] = 0
				slot++
				state = "B"
			}
		case "E":
			if set[slot] == 0 {
				set[slot] = 1
				slot--
				state = "F"
			} else {
				set[slot] = 1
				slot--
				state = "C"
			}
		case "F":
			if set[slot] == 0 {
				set[slot] = 1
				slot++
				state = "D"
			} else {
				set[slot] = 1
				slot++
				state = "A"
			}

		}
	}
	count := 0
	for _, val := range set {
		count += val
	}
	fmt.Println(count)
	//805298 too high
}
