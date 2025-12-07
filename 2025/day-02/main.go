package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/roessland/advent-of-code/2025/aocutil"
)

type Interval struct {
	A, B int
}

func ReadInput() []Interval {
	f, err := os.Open("input.txt")
	// f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	line, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	nums := aocutil.GetNumsInString[int](string(line))
	input := make([]Interval, 0, len(nums)/2)
	for i := 0; i < len(nums); i += 2 {
		input = append(input, Interval{nums[i], nums[i+1]})
	}
	return input
}

func Part1() {
	input := ReadInput()
	fmt.Println(input)
	sum_of_invalid_ids := 0
	for _, pair := range input {
		// Sum invalid IDs in the inclusive range formed by pair.
		start := pair.A
		end := pair.B
		for i := start; i <= end; i++ {
			base10 := strconv.Itoa(i)
			base10len := len(base10)
			found_one := false
			for repetee_len := 1; repetee_len <= (base10len / 2); repetee_len++ {
				if (base10len % repetee_len) != 0 {
					// String length not divisible by repetee length cannot be
					// made up of repretitions of repetee.
					continue
				}
				repetee := base10[:repetee_len]
				if strings.Repeat(repetee, base10len/repetee_len) == base10 {
					// ID is a number repeated l times
					found_one = true
					break
				}
			}
			if found_one {
				sum_of_invalid_ids += i
			}
		}
	}
	fmt.Println(sum_of_invalid_ids)
	fmt.Println(time.Now())
}

func main() {
	Part1()
}
