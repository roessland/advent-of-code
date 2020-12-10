package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Ways(nums []int) int {
	cache := map[int]int{len(nums)-1: 1}
	var w func(int)int
	w = func(i0 int)int {
		cached, ok := cache[i0]
		if ok {
			return cached
		}
		n0 := nums[i0]
		ways := 0
		for i := i0+1;  i < len(nums) && nums[i] <= n0 + 3; i++ {
			ways += w(i)
		}
		cache[i0] = ways
		return ways
	}
	return w(0)
}

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

	fmt.Println(Ways(nums))
}
