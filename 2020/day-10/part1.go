package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)


func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	nums := []int{0}
	max := 0
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			panic("nope")
		}
		nums = append(nums ,n)
		if n > max {
			max = n
		}
	}
	nums = append(nums, max+3)
	sort.Ints(nums)

	diffs := map[int]int{}
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		diffs[diff]++
	}
	fmt.Println(diffs[1]*diffs[3])
}
