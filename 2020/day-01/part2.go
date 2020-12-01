package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func sortedListContains(nums []int, n int) bool {
	insertIdx := sort.SearchInts(nums, n)
	if insertIdx > 0 && insertIdx < len(nums) && nums[insertIdx] == n {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var nums []int
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)

	out:
	for _, n1 := range nums {
		for _, n2 := range nums {
			n3 := 2020-n1-n2
			// search for n3
			if sortedListContains(nums, n3) {
				fmt.Println(n1*n2*n3)
				break out
			}
		}
	}
}