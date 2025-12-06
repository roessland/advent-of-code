// Package day03 solves AoC 2025 Day 3
package day03

import (
	"embed"
	"fmt"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var InputFile embed.FS

func ReadInput(inputName string) (banks [][]int) {
	lines := aocutil.ReadLinesAsBytes(inputName)

	for _, line := range lines {
		var bank []int
		for _, c := range line {
			bank = append(bank, int(c)-'0')
		}
		banks = append(banks, bank)
	}
	return banks
}

// MaxIdx returns the index in [a, b) pointing to the largest element in the array.
func MaxIdx(ns []int, a, b int) (iMax int) {
	iMax = a
	nMax := ns[a]
	for i := iMax; i < b; i++ {
		if ns[i] > nMax {
			iMax = i
			nMax = ns[i]
		}
	}
	return iMax
}

func LargestJoltPossible2(bank []int) int {
	i1 := MaxIdx(bank, 0, len(bank)-1)
	i2 := MaxIdx(bank, i1+1, len(bank))
	j1, j2 := bank[i1], bank[i2]
	return 10*j1 + j2
}

func TotalOutputJoltage2(banks [][]int) (total int) {
	for _, bank := range banks {
		total += LargestJoltPossible2(bank)
	}
	return total
}

func Part1(banks [][]int) (total int) {
	for _, bank := range banks {
		total += LargestJoltPossible2(bank)
	}
	return total
}

func LargestJoltPossibleN(bank []int, K int) int {
	is := []int{}
	for k := range K {
		var i int
		if k == 0 {
			i = MaxIdx(bank, 0, len(bank)-K+k+1)
		} else {
			i = MaxIdx(bank, is[k-1]+1, len(bank)-K+k+1)
		}
		is = append(is, i)
	}
	fmt.Println(IdxsToJoltage(bank, is))
	return IdxsToJoltage(bank, is)
}

func IdxsToJoltage(bank []int, is []int) int {
	joltage := 0
	for pos := range is {
		joltage = joltage*10 + bank[is[pos]]
	}
	return joltage
}

func Part2(banks [][]int) (total int) {
	for _, bank := range banks {
		total += LargestJoltPossibleN(bank, 12)
	}
	return total
}
