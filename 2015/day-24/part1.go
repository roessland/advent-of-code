package main

import "fmt"
import "math/rand"

func Prod(part []int) int {
	prod := 1
	for _, weight := range part {
		prod *= weight
	}
	return prod
}

func Sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}

func main() {
	ps := []int{1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59,
		61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}

	// Each partition must weight 520
	targetSum := Sum(ps) / 3

	minProd := 99999999999999999
	N := 6
	p := make([]int, N)
	for {
		perm := rand.Perm(len(ps))[:N]
		for i, j := range perm {
			p[i] = ps[j]
		}
		if Sum(p) == targetSum {
			prodP := Prod(p)
			if prodP < minProd {
				minProd = prodP
				fmt.Printf("%v\n", minProd)
			}
		}
	}
}
