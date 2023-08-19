package aocutil

import (
	"github.com/davecgh/go-spew/spew"
	"regexp"
	"strconv"
	"strings"
)

// on x=-69723..-64530,y=22399..28572,z=-47850..-23758

func GetIntsInString(line string) []int {
	return GetNumsInString[int](line)
}

func GetNumsInString[N ~int](line string) []N {
	var lineRe = regexp.MustCompile(`(^|[^a-zA-Z])(-?\d+)`)
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
