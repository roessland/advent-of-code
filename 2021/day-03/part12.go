package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(bins []string) {
	N := len(bins)
	K := len(bins[0])
	oneCounts := make([]int, K)
	for _, bin := range bins {
		for i, bit := range bin {
			if bit == '1' {
				oneCounts[i]++
			}
		}
	}

	gammaRateStr := make([]byte, K)
	epsilonStr := make([]byte, K)
	for k := range gammaRateStr {
		if oneCounts[k] > N/2 {
			gammaRateStr[k] = '1'
			epsilonStr[k] = '0'
		} else {
			gammaRateStr[k] = '0'
			epsilonStr[k] = '1'
		}
	}

	gammaRate, err := strconv.ParseUint(string(gammaRateStr), 2, 64)
	if err != nil {
		panic(fmt.Sprint(gammaRate))
	}
	epsilon, err := strconv.ParseUint(string(epsilonStr), 2, 64)
	if err != nil {
		panic("nah2")
	}

	fmt.Println(gammaRate*epsilon)
}

func find(bins []string, k int, flip bool) int {
	if len(bins) == 1 {
		rating, err := strconv.ParseUint(bins[0], 2, 64)
		if err != nil {
			panic("whoops")
		}
		return int(rating)
	}

	var oneCount int
	for _, bin := range bins {
		if bin[k] == '1' {
			oneCount++
		}
	}

	var mostCommon byte
	criteria := oneCount >= len(bins) - oneCount
	if flip {
		criteria = !criteria
	}
	if criteria {
		mostCommon = '1'
	} else {
		mostCommon = '0'
	}

	var filteredBins []string
	for _, bin := range bins {
		if bin[k] == mostCommon {
			filteredBins = append(filteredBins, bin)
		}
	}
	return find(filteredBins, k+1, flip)
}

func part2(bins []string) {
	fmt.Println(find(bins, 0, false) * find(bins, 0, true))
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var cmds []string
	for scanner.Scan() {
		line := scanner.Text()
		cmds = append(cmds, line)
	}

	part1(cmds)
	part2(cmds)
}
