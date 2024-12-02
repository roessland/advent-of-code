package day02

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

//go:embed input*.txt
var Input embed.FS

type Level = int

func ReadInput(inputName string) [][]Level {
	return aocutil.ReadFileAsInts(Input, inputName)
}

func IsAllIncreasing(report []Level) bool {
	for i := 1; i < len(report); i++ {
		if report[i] <= report[i-1] {
			return false
		}
	}
	return true
}

func IsAllDecreasing(report []Level) bool {
	for i := 1; i < len(report); i++ {
		if report[i] >= report[i-1] {
			return false
		}
	}
	return true
}

func HasValidDifferences(report []Level) bool {
	for i := 1; i < len(report); i++ {
		absDiff := mathutil.AbsInt(report[i] - report[i-1])
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func IsSafe(report []Level) bool {
	return (IsAllIncreasing(report) || IsAllDecreasing(report)) && HasValidDifferences(report)
}

func Part1(inputName string) int {
	reports := ReadInput(inputName)
	numSafe := 0
	for i := range reports {
		if IsSafe(reports[i]) {
			numSafe++
		}
	}
	return numSafe
}

func DampenProblem(report []Level, idx int) []Level {
	dampenedLevels := make([]Level, len(report)-1)
	copy(dampenedLevels, report[:idx])
	copy(dampenedLevels[idx:], report[idx+1:])
	return dampenedLevels
}

func Part2(inputName string, anim *AnimationHooks) int {
	anim = orNoopHooks(anim)

	reports := ReadInput(inputName)
	anim.SetLength(len(reports))

	numSafe := 0
nextReport:
	for i := range reports {
		anim.IncreaseColor(i)
		if IsSafe(reports[i]) {
			numSafe++
			anim.SetColor(i, 1)
			continue nextReport
		}
		for j := range reports[i] {
			anim.IncreaseColor(i)
			if IsSafe(DampenProblem(reports[i], j)) {
				numSafe++
				anim.SetColor(i, 1)
				continue nextReport
			}
		}
		anim.SetColor(i, 0)
	}
	return numSafe
}
