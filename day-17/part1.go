package main

import "fmt"
import "sort"

var depth_frequency map[int]float64

func Factorial(n int) float64 {
	f := 1.0
	for i := 2; i <= n; i++ {
		f *= float64(i)
	}
	return f
}

func C(cs []int, total int, depth int) float64 {
	if total < 0 {
		return 0.0
	} else if total == 0 {
		depth_frequency[depth] += 1.0 / Factorial(depth)
		return 1.0 / Factorial(depth)
	}
	ways := 0.0
	for i := 0; i < len(cs) && total-cs[i] >= 0; i++ {
		newcs := make([]int, 0, len(cs)-1)
		newcs = append(newcs, cs[:i]...)
		newcs = append(newcs, cs[i+1:]...)
		ways += C(newcs, total-cs[i], depth+1)
	}
	return ways
}

func main() {
	depth_frequency = make(map[int]float64)
	containers := []int{50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40}
	sort.Ints(containers)
	fmt.Printf("Ways: %v\n", C(containers, 150, 0))
	fmt.Printf("Depth frequencies: %v\n", depth_frequency)
}
