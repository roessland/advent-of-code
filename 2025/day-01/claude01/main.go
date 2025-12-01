package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (byte, int) {
	dir := line[0]
	dist, _ := strconv.Atoi(line[1:])
	return dir, dist
}

func rotate(pos int, dir byte, dist int) int {
	if dir == 'L' {
		pos -= dist
	} else {
		pos += dist
	}
	// Handle modulo for negative numbers
	pos = ((pos % 100) + 100) % 100
	return pos
}

func solve(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pos := 50
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		dir, dist := parseLine(line)
		pos = rotate(pos, dir, dist)
		if pos == 0 {
			count++
		}
	}

	return count
}

// countZeroCrossings counts how many times the dial passes through 0
// during a rotation from pos in direction dir by dist clicks.
func countZeroCrossings(pos int, dir byte, dist int) int {
	var firstHit int
	if dir == 'L' {
		// Going left (decreasing): we hit 0 at step pos, pos+100, pos+200, ...
		// If pos == 0, first hit is at step 100
		if pos == 0 {
			firstHit = 100
		} else {
			firstHit = pos
		}
	} else {
		// Going right (increasing): we hit 0 at step (100-pos), (100-pos)+100, ...
		// If pos == 0, first hit is at step 100
		if pos == 0 {
			firstHit = 100
		} else {
			firstHit = 100 - pos
		}
	}

	if firstHit > dist {
		return 0
	}
	// Count: firstHit, firstHit+100, firstHit+200, ... all <= dist
	return (dist-firstHit)/100 + 1
}

func solve2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pos := 50
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		dir, dist := parseLine(line)
		count += countZeroCrossings(pos, dir, dist)
		pos = rotate(pos, dir, dist)
	}

	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Part 1:", solve(string(data)))
	fmt.Println("Part 2:", solve2(string(data)))
}
