package main

import "fmt"
import "math/rand"

func Sums(parts [][]int) []int {
	sums := make([]int, len(parts))
	for i, _ := range parts {
		for j, _ := range parts[i] {
			sums[i] += parts[i][j]
		}
	}
	return sums
}

func SmallestPart(parts [][]int) []int {
	smallest := parts[0]
	if len(parts[1]) < len(smallest) {
		smallest = parts[1]
	}
	if len(parts[2]) < len(smallest) {
		smallest = parts[2]
	}
	return smallest
}

func QE(part []int) int {
	prod := 1
	for _, weight := range part {
		prod *= weight
	}
	return prod
}

func main() {
	minPackages := 99999
	//minQE := 99999
	// Create random partitions
	for {
		ps := []int{1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59,
			61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
		parts := make([][]int, 3)
		for len(ps) > 0 {
			i := rand.Intn(3)
			j := rand.Intn(len(ps))
			parts[i] = append(parts[i], ps[j])
			ps = append(ps[:j], ps[j+1:]...)
		}
		sums := Sums(parts)
		if sums[0] == sums[1] && sums[1] == sums[2] {
			smallest := SmallestPart(parts)
			if len(smallest) <= minPackages {
				minPackages = len(smallest)
				fmt.Printf("Candidate with %v packages and QE %v.\n", minPackages, QE(smallest))
			}
			fmt.Printf("%v\n", sums)
		}
	}
}
