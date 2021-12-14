package main

import (
	"bufio"
	"fmt"
	. "github.com/roessland/gopkg/mathutil"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()
	polymer, rules := ReadInput()
	part12(polymer, rules)
	fmt.Println(time.Since(t0)) // 1 ms
}

func part12(polymer string, rules map[[2]byte]byte) {
	// Main insight: No need to store the order of elements in the polymer,
	// only the count of each pair.
	pairCounts := map[[2]byte]int64{}
	for i := 1; i < len(polymer); i++ {
		pairCounts[[2]byte{polymer[i-1], polymer[i]}]++
	}

	var i = 0
	for ; i < 10; i++ {
		pairCounts = applyRules(pairCounts, rules)
	}
	fmt.Println("Part 1:", score(pairCounts))

	for ; i < 40; i++ {
		pairCounts = applyRules(pairCounts, rules)
	}
	fmt.Println("Part 2:", score(pairCounts))
}

func applyRules(prevPairCounts map[[2]byte]int64, rules map[[2]byte]byte) map[[2]byte]int64 {
	pairCounts := make(map[[2]byte]int64)
	for prevPair, prevCount := range prevPairCounts {
		between := rules[prevPair]
		pairCounts[[2]byte{prevPair[0], between}] += prevCount
		pairCounts[[2]byte{between, prevPair[1]}] += prevCount
	}
	return pairCounts
}

func score(pairCounts map[[2]byte]int64) int64 {
	elementCounts := make(map[byte]int64)
	for pair, pairCount := range pairCounts {
		elementCounts[pair[0]] += pairCount
		elementCounts[pair[1]] += pairCount
	}
	var min, max int64 = math.MaxInt64, 0
	for element := range elementCounts {
		// Fix count since we have counted things twice (except those at edges)
		count := (elementCounts[element] + 1) / 2
		min, max = MinInt64(min, count), MaxInt64(max, count)
	}
	return max - min
}

func ReadInput() (string, map[[2]byte]byte) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	rules := make(map[[2]byte]byte)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		rules[[2]byte{parts[0][0], parts[0][1]}] = parts[1][0]
	}
	return polymer, rules
}
