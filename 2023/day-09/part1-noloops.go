package main

import (
	"fmt"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	inputLines := aocutil.ReadLines("input.txt")
	histories := mapStringsToIntArrays(inputLines, func(str string) []int {
		return parseLine(str)
	})

	extrapolatedVals := mapIntsToInt(histories, func(hist []int) int {
		fmt.Println("Mapping history", hist)
		return getNextValue(hist)
	})

	sum := reduce(extrapolatedVals, add)

	fmt.Println("Sum:", sum)
}

func reduce(arr []int, f func(a, b int) int) int {
	if len(arr) == 0 {
		panic("Cannot reduce empty array")
	}
	if len(arr) == 1 {
		return arr[0]
	}
	rest, tail := arr[:len(arr)-1], arr[len(arr)-1]
	return f(reduce(rest, f), tail)
}

func add(a, b int) int {
	return a + b
}

func getNextValue(history []int) int {
	if areAllZero(history) {
		return 0
	}
	return history[len(history)-1] + getNextValue(getDiffs(history))
}

func getDiffs(history []int) []int {
	if len(history) < 2 {
		return nil
	}
	rest, tail := history[:len(history)-1], history[len(history)-1]
	return append(getDiffs(rest), tail-rest[len(rest)-1])
}

func areAllZero(history []int) bool {
	if len(history) == 0 {
		return true
	}
	return history[0] == 0 && areAllZero(history[1:])
}

func parseLine(line string) []int {
	return aocutil.GetIntsInString(line)
}

func mapStringsToIntArrays(ss []string, f func(str string) []int) [][]int {
	if len(ss) == 0 {
		return nil
	}
	rest, tail := ss[:len(ss)-1], ss[len(ss)-1]
	return append(mapStringsToIntArrays(rest, f), f(tail))
}

func mapIntsToInt(arrs [][]int, f func([]int) int) []int {
	if len(arrs) == 0 {
		return nil
	}
	rest, tail := arrs[:len(arrs)-1], arrs[len(arrs)-1]
	return append(mapIntsToInt(rest, f), f(tail))
}
