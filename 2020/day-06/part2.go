package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	reader := bufio.NewScanner(f)
	sum := 0
	yes := map[rune]int{}
	numPeople := 0
	for reader.Scan() { // Remember double newline at end of input
		qs := reader.Text()
		if qs == "" {
			allYes := 0
			for _, y := range yes {
				if y == numPeople {
					allYes++
				}
			}
			sum += allYes
			numPeople = 0
			yes = map[rune]int{}
			continue
		}
		numPeople++
		for _, q := range qs {
			yes[q]++
		}
	}
	fmt.Println(sum)
}