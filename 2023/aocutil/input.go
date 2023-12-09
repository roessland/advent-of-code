package aocutil

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// on x=-69723..-64530,y=22399..28572,z=-47850..-23758

func GetIntsInString(line string) []int {
	return GetNumsInString[int](line)
}

func GetNumsInString[N ~int](line string) []N {
	lineRe := regexp.MustCompile(`(^|[^a-zA-Z])(-?\d+)`)
	matches := lineRe.FindAllStringSubmatch(line, -1)
	nums := make([]N, 0, len(matches)-1)
	for _, match := range matches {
		n, err := strconv.Atoi(strings.TrimSpace(match[2]))
		if err != nil {
			spew.Dump(matches)
			panic(err)
		}
		nums = append(nums, N(n))
	}
	return nums
}

// ReadFile wraps string(os.ReadFile) and panics on error.
// Also removes trailing newline.
func ReadFile(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(f), "\n")
}
