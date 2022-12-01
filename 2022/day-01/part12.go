package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

func main() {
	nums := ReadInput()
	part1(nums)
	part2(nums)
}

func part1(elves [][]int) {
	maxTotalCals := 0
	for _, elf := range elves {
		totalCals := 0
		for _, cals := range elf {
			totalCals += cals
		}
		if totalCals > maxTotalCals {
			maxTotalCals = totalCals
		}
	}
	fmt.Println(maxTotalCals)
}

func part2(elves [][]int) {
	totalCals := make([]int, len(elves))
	for i, elf := range elves {
		for _, cals := range elf {
			totalCals[i] += cals
		}
	}

	sort.Ints(totalCals)

	sum := 0
	sum += totalCals[len(totalCals)-1]
	sum += totalCals[len(totalCals)-2]
	sum += totalCals[len(totalCals)-3]
	fmt.Println(sum)
}

func ReadInput() [][]int {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var nums [][]int
	nums = append(nums, []int{})

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			nums = append(nums, []int{})
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic("nope")
			}
			nums[len(nums)-1] = append(nums[len(nums)-1], n)
		}
	}
	return nums
}
