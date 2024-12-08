package day07

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Eq struct {
	Numbers           []int
	CalibrationResult int
}

func ReadInput(inputName string) []Eq {
	eqs := make([]Eq, 0)
	in := aocutil.FSGetIntsInStringLines(Input, inputName)
	for _, line := range in {
		eqs = append(eqs, Eq{
			CalibrationResult: line[0],
			Numbers:           line[1:],
		})
	}
	return eqs
}

func CanBeTrue1(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}

	original := nums[1]
	defer func() { nums[1] = original }()

	add := nums[0] + nums[1]
	mul := nums[0] * nums[1]

	nums[1] = mul
	if CanBeTrue1(nums[1:], target) {
		return true
	}

	nums[1] = add
	return CanBeTrue1(nums[1:], target)
}

func CanBeTrue2(nums []int, target int, d int) bool {
	if len(nums) == 0 || nums[0] > target {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}

	original := nums[1]
	add := nums[0] + nums[1]
	mul := nums[0] * nums[1]
	cat := Cat(nums[0], nums[1])

	nums[1] = mul
	if CanBeTrue2(nums[1:], target, d+1) {
		return true
	}

	nums[1] = add
	if CanBeTrue2(nums[1:], target, d+1) {
		return true
	}

	nums[1] = cat
	if CanBeTrue2(nums[1:], target, d+1) {
		return true
	}

	nums[1] = original
	return false
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
	input := ReadInput(inputName)
	sum1, sum2 := 0, 0
	for _, eq := range input {
		if CanBeTrue1(eq.Numbers, eq.CalibrationResult) {
			sum1 += eq.CalibrationResult
		}
		if CanBeTrue2(eq.Numbers, eq.CalibrationResult, 0) {
			sum2 += eq.CalibrationResult
		}
	}
	return sum1, sum2
}
