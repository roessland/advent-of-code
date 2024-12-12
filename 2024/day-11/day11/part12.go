package day11

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func Next(arr []int) []int {
	next := []int{}
	for _, before := range arr {
		if before == 0 {
			next = append(next, 1)
			continue
		}

		numDigits := NumDigits2(before)
		if numDigits%2 == 0 {
			a, b := Split(before, numDigits)
			next = append(next, a, b)
			continue
		}

		next = append(next, before*2024)
	}
	return next
}

func NumDigits2(n int) int {
	i := 1
	for n > 9 {
		n /= 10
		i++
	}
	return i
}

func Split(n, numDigits int) (int, int) {
	t := Pow10(numDigits / 2)
	return n / t, n % t
}

func Pow10(n int) int {
	p := 1
	for n > 0 {
		p *= 10
		n--
	}
	return p
}

var cache = map[[2]int]int{}

func N(n int, its int) int {
	if its == 0 {
		return 1
	}

	key := [2]int{n, its}
	cached, ok := cache[key]
	if ok {
		return cached
	}

	if n == 0 {
		return N(1, its-1)
	}

	numDigits := NumDigits2(n)
	if numDigits%2 == 0 {
		a, b := Split(n, numDigits)
		cache[key] = N(a, its-1) + N(b, its-1)
		return cache[key]
	}

	cache[key] = N(n*2024, its-1)
	return cache[key]
}

func Part12(inputName string) (int, int) {
	arr := aocutil.FSGetIntsInStringLines(Input, inputName)[0]

	sum1 := 0
	for _, n := range arr {
		sum1 += N(n, 25)
	}

	sum2 := 0
	for _, n := range arr {
		sum2 += N(n, 75)
	}

	return sum1, sum2
}
