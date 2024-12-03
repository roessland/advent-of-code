package day03

import (
	"embed"
	"regexp"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func Part12(inputName string) (int, int) {
	mem := aocutil.FSReadFile(Input, inputName)
	re := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`)
	matches := re.FindAllStringSubmatch(mem, -1)
	sum1 := 0
	sum2 := 0
	enabled := 1
	for _, match := range matches {
		if match[1] == "mul" {
			a := aocutil.Atoi(match[2])
			b := aocutil.Atoi(match[3])
			sum1 += a * b
			sum2 += a * b * enabled
		} else if match[0] == "do()" {
			enabled = 1
		} else if match[0] == "don't()" {
			enabled = 0
		}
	}
	return sum1, sum2
}
