package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := aocutil.ReadLines("input.txt")
	times := aocutil.GetIntsInString(lines[0])
	dists := aocutil.GetIntsInString(lines[1])

	prod := 1
	for i := range times {
		time := times[i]
		dist := dists[i]
		ways := numWays(time, dist)
		prod *= ways
	}
	fmt.Println(prod)
}

func part2() {
	lines := aocutil.ReadLines("input.txt")
	timesStr := lines[0]
	timesStr = strings.ReplaceAll(timesStr, "Time:", "")
	timesStr = strings.ReplaceAll(timesStr, " ", "")
	time := aocutil.Atoi(timesStr)

	distStr := lines[1]
	distStr = strings.ReplaceAll(distStr, "Distance:", "")
	distStr = strings.ReplaceAll(distStr, " ", "")
	recordDist := aocutil.Atoi(distStr)

	fmt.Println(numWays(time, recordDist))
}

func numWays(time, recordDist int) int {
	ways := 0
	for hold := 0; hold <= time; hold++ {
		if hold*(time-hold) > recordDist {
			ways++
		}
	}
	return ways
}
