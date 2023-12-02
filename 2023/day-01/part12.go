package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lines := ReadInput()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	digitRegex := regexp.MustCompile(`\d`)
	for _, line := range lines {
		matches := digitRegex.FindAllString(line, -1)
		fst, err := strconv.Atoi(matches[0])
		check(err)
		lst, err := strconv.Atoi(matches[len(matches)-1])
		check(err)
		val := 10*fst + lst
		sum += val
	}
	fmt.Println(sum)
}

func toInt(s string) int {
	switch s {
	case "1":
		fallthrough
	case "one":
		return 1
	case "2":
		fallthrough
	case "two":
		return 2
	case "3":
		fallthrough
	case "three":
		return 3
	case "4":
		fallthrough
	case "four":
		return 4
	case "5":
		fallthrough
	case "five":
		return 5
	case "6":
		fallthrough
	case "six":
		return 6
	case "7":
		fallthrough
	case "seven":
		return 7
	case "8":
		fallthrough
	case "eight":
		return 8
	case "9":
		fallthrough
	case "nine":
		return 9
	}
	panic(fmt.Sprintf("unknown digit '%s'", s))
}

type Match struct {
	Str   string
	Index int
}

// Go regexes don't support overlapping matches, and fails to find
// "one" in "twone", and "nine" in "sevenine".
func getMatches(line string) []string {
	nums := []string{
		"1", "one",
		"2", "two",
		"3", "three",
		"4", "four",
		"5", "five",
		"6", "six",
		"7", "seven",
		"8", "eight",
		"9", "nine",
	}

	matches := []Match{}
	for _, num := range nums {
		idx0 := strings.Index(line, num)
		if idx0 == -1 {
			continue
		}
		matches = append(matches, Match{Index: idx0, Str: num})
		idx1 := strings.LastIndex(line, num)
		if idx1 != idx0 {
			matches = append(matches, Match{Index: idx1, Str: num})
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Index < matches[j].Index
	})

	matchStrings := []string{}
	for _, match := range matches {
		matchStrings = append(matchStrings, match.Str)
	}

	return matchStrings
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		matches := getMatches(line)
		fst := toInt(matches[0])
		lst := toInt(matches[len(matches)-1])
		val := 10*fst + lst
		sum += val
	}
	fmt.Println(sum)
}

func ReadInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines
}
