package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(id int64) bool {
	s := strconv.FormatInt(id, 10)
	n := len(s)

	// Check all possible pattern lengths (1 to n/2)
	// Pattern must repeat at least 2 times
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue // Pattern length must divide total length evenly
		}

		pattern := s[:patternLen]
		isRepeat := true
		for i := patternLen; i < n; i += patternLen {
			if s[i:i+patternLen] != pattern {
				isRepeat = false
				break
			}
		}

		if isRepeat {
			return true
		}
	}

	return false
}

func sumInvalidIDsInRange(start, end int64) int64 {
	var sum int64
	for id := start; id <= end; id++ {
		if isInvalidID(id) {
			sum += id
		}
	}
	return sum
}

func solve(input string) int64 {
	input = strings.TrimSpace(input)
	// Remove trailing comma if present
	input = strings.TrimSuffix(input, ",")

	ranges := strings.Split(input, ",")
	var total int64

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		total += sumInvalidIDsInRange(start, end)
	}

	return total
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result := solve(string(data))
	fmt.Println(result)
}
