package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Pool struct {
	Size int
	Nums []int
}

func (p *Pool) Push(num int) {
	p.Nums = append([]int{num}, p.Nums...)
}

func (p *Pool) Pop() int {
	ret := p.Nums[len(p.Nums)-1]
	p.Nums = p.Nums[0:len(p.Nums)-1]
	return ret
}

func (p *Pool) Add(num int) {
	if len(p.Nums) != p.Size {
		panic("use Push to fill preamble")
	}
	p.Pop()
	p.Push(num)
}

func (p *Pool) Valid(num int) bool {
	for i := 0; i < len(p.Nums); i++ {
		for j := 0; j < len(p.Nums); j++ {
			if i == j {
				continue
			}
			if p.Nums[i] + p.Nums[j] == num {
				return true
			}
		}
	}
	return false
}

func FindEncryptionWeakness(nums []int, invalid int) int {
	prefixsum := make([]int, len(nums)+1)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixsum[i+1] = sum
	}
	var firstIdx, lastIdx int
	for i := 0; i < len(prefixsum)-2; i++ {
		for j := i+2; j < len(prefixsum); j++ {
			if prefixsum[j]-prefixsum[i] == invalid {
				firstIdx, lastIdx = i, j-1
			}
		}
	}

	min, max := math.MaxInt32, 0
	for i := firstIdx; i <= lastIdx; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return min + max
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

	p := Pool{Size: 25}
	var invalid int
	for i, num := range nums {
		if i < p.Size {
			p.Push(num)
			continue
		}
		if p.Valid(num) {
			p.Add(num)
		} else {
			invalid = num
			fmt.Println("Part 1: ", invalid)
			break
		}
	}

	fmt.Println("Part 2: ", FindEncryptionWeakness(nums, invalid))
}