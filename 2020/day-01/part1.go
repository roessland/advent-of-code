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
			if n1+n2 == 2020 {
				fmt.Println(n1*n2)
				break out
			}
		}
	}
}