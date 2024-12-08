package day07

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func CanBeTrue(result int, fst int, rest []int, canCat bool) bool {
	if fst > result { // Around 25% speedup
		return false
	}
	if len(rest) == 0 {
		return fst == result
	}

	return CanBeTrue(result, fst+rest[0], rest[1:], canCat) ||
		CanBeTrue(result, fst*rest[0], rest[1:], canCat) ||
		(canCat && CanBeTrue(result, Cat(fst, rest[0]), rest[1:], canCat))
}

func Cat(a, b int) int {
	return 10*a*Order(b) + b
}

func Order(b int) int {
	n := 1
	for b >= 10 {
		b /= 10
		n *= 10
	}
	return n
}

func Part12(inputName string) (int, int) {
	input := aocutil.FSGetIntsInStringLines(Input, inputName)
	sum1, sum2 := 0, 0
	for _, line := range input {
		res, fst, rest := line[0], line[1], line[2:]

		if CanBeTrue(res, fst, rest, false) {
			sum1 += res
		}
		if CanBeTrue(res, fst, rest, true) {
			sum2 += res
		}
	}
	return sum1, sum2
}
