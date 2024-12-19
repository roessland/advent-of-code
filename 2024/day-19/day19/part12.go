package day19

import (
	"embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

func ReadInput(inputName string) (patterns, designs []string) {
	lines := strings.Split(aocutil.FSReadFile(Input, inputName), "\n")
	return strings.Split(lines[0], ", "), lines[2:]
}

func Part1(patterns, designs []string) int {
	pat := fmt.Sprintf(`^(%s)+$`, strings.Join(patterns, "|"))
	re := regexp.MustCompile(pat)

	numPossible := 0
	for _, design := range designs {
		if re.MatchString(design) {
			numPossible++
		}
	}
	return numPossible
}

func Ways(patterns []string, design string, start int) int {
	cache := map[int]int{}

	var ways func(design string, start int) int
	ways = func(design string, start int) int {
		if cached, ok := cache[start]; ok {
			return cached
		}

		if len(design[start:]) == 0 {
			return 1
		}

		numWays := 0
		for _, pattern := range patterns {
			if strings.HasPrefix(design[start:], pattern) {
				numWays += ways(design, start+len(pattern))
			}
		}
		cache[start] = numWays
		return numWays
	}
	return ways(design, start)
}

func Part2(patterns, designs []string) int {
	sum := 0
	for _, design := range designs {
		sum += Ways(patterns, design, 0)
	}
	return sum
}

func Part12(inputName string) (int, int) {
	patterns, designs := ReadInput(inputName)
	return Part1(patterns, designs), Part2(patterns, designs)
}
