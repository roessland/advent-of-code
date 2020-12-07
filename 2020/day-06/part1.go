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
	yes := map[rune]bool{}
	for reader.Scan() { // Remember double newline at end of input
		qs := reader.Text()
		if qs == "" {
			sum += len(yes)
			yes = map[rune]bool{}
			continue
		}
		for _, q := range qs {
			yes[q] = true
		}
	}
	fmt.Println(sum)
}