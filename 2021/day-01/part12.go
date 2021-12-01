package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(nums []int) {
	incs := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			incs++
		}
	}

	fmt.Println("Part 1: ", incs)
}

func part2(nums []int) {
	incs := 0
	for i := 3; i < len(nums); i++ {
		if  nums[i] > nums[i-3]  {
			incs++
		}
	}
	fmt.Println("Part 2: ", incs)

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
			panic("nope")
		}
		nums = append(nums, n)
	}

	part1(nums)
	part2(nums)
}
