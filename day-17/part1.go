package main

import "fmt"

func C(cs []int, total int) int {
	if total < 0 {
		return 0
	} else if total == 0 {
		return 1
	}
	ways := 0
	for i := 0; i < len(cs); i++ {
		newcs := make([]int, 0, len(cs)-1)
		newcs = append(newcs, cs[:i]...)
		newcs = append(newcs, cs[i+1:]...)
		ways += C(newcs, total-cs[i])
	}
	return ways
}

func main() {
	containers := []int{50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40}
	fmt.Printf("Ways: %v\n", C(containers, 150))
}
