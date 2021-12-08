package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func Dist(crabs []int, n int) (sumDist int) {
	for _, crab := range crabs {
		sumDist += mathutil.AbsInt(crab - n)
	}
	return sumDist
}

func Dist2(crabs []int, n int) (sumDist2 int) {
	for _, crab := range crabs {
		dist1 := mathutil.AbsInt(crab - n)
		sumDist2 += dist1 * (dist1 + 1) / 2
	}
	return sumDist2
}

func MinMax(crabs []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32
	for _, crab := range crabs {
		min = mathutil.MinInt(min, crab)
		max = mathutil.MaxInt(max, crab)
	}
	return min, max
}

func main() {
	crabs := ReadInput()

	minDist, minDist2 := math.MaxInt32, math.MaxInt32
	min, max := MinMax(crabs)
	for pos := min; pos <= max; pos++ {
		dist := Dist(crabs, pos)
		if dist < minDist {
			minDist = dist
		}

		dist2 := Dist2(crabs, pos)
		if dist2 < minDist2 {
			minDist2 = dist2
		}
	}

	fmt.Println("Part 1:", minDist)
	fmt.Println("Part 2:", minDist2)
}

func ReadInput() []int {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	var crabs []int
	for _, str := range strings.Split(string(buf), ",") {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, n)
	}
	return crabs
}
