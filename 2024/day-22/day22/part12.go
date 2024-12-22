package day22

import (
	"embed"
	"iter"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func ReadInput(inputName string) (secrets []int) {
	lines := aocutil.FSGetIntsInStringLines(Input, inputName)
	nums := make([]int, len(lines))
	for i, line := range lines {
		nums[i] = line[0]
	}
	return nums
}

func Next(secret int) int {
	secret ^= secret * 64
	secret %= 16777216

	secret ^= secret / 32
	secret %= 16777216

	secret ^= secret * 2048
	secret %= 16777216
	return secret
}

func Nth(secret int, n int) int {
	for i := 0; i < n; i++ {
		secret = Next(secret)
	}
	return secret
}

type FourConsecutiveChanges struct {
	Changes [4]int8
	Price   int8
}

func EachChange(secret int) iter.Seq[FourConsecutiveChanges] {
	A := secret
	B := Next(A)
	C := Next(B)
	D := Next(C)
	E := Next(D)
	return func(yield func(FourConsecutiveChanges) bool) {
		for i := 0; i < 2000-3; i++ {
			ab, bc, cd, de := (B%10)-(A%10), (C%10)-(B%10), (D%10)-(C%10), (E%10)-(D%10)
			changes := FourConsecutiveChanges{
				Changes: [4]int8{int8(ab), int8(bc), int8(cd), int8(de)},
				Price:   int8(E % 10),
			}
			if !yield(FourConsecutiveChanges(changes)) {
				return
			}
			A, B, C, D, E = B, C, D, E, Next(E)
		}
	}
}

func Part1(secrets []int) int {
	sum := 0
	for _, secret := range secrets {
		sum += Nth(secret, 2000)
	}
	return sum
}

func Bananas(secret int) map[[4]int8]int {
	bananas := make(map[[4]int8]int)
	for change := range EachChange(secret) {
		if _, ok := bananas[change.Changes]; !ok {
			bananas[change.Changes] += int(change.Price)
		}
	}
	return bananas
}

func Part2(secrets []int) int {
	bananas := map[[4]int8]int{}

	for _, secret := range secrets {
		buyerBananas := Bananas(secret)
		for k, v := range buyerBananas {
			bananas[k] += v
		}
	}

	maxBananas := 0
	for _, v := range bananas {
		if v > maxBananas {
			maxBananas = v
		}
	}
	return maxBananas
}

func Part12(inputName string) (int, int) {
	secrets := ReadInput(inputName)
	return Part1(secrets), Part2(secrets)
}
