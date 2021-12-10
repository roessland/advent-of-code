package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func push(stack []rune, c rune) []rune {
	return append(stack, c)
}

func pop(stack []rune) ([]rune, rune) {
	if len(stack) == 0 {
		panic("nah")
	}
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func peek(stack []rune) rune {
	if len(stack) == 0 {
		panic("nah")
	}
	return stack[len(stack)-1]
}

var isClosingBrace = map[rune]bool{
	']': true,
	'}': true,
	')': true,
	'>': true,
}

func part1(lines []string) []string {
	closerFor := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
		'<': '>',
	}
	scoreFor := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	totalScore := 0
	var stack []rune
	var incompleteLines []string
	for _, line := range lines {
		var broken = false
	corrupted:
		for _, c := range line {
			if len(stack) == 0 {
				stack = push(stack, c)
			} else if isClosingBrace[c] {
				if closerFor[peek(stack)] != c {
					totalScore += scoreFor[c]
					broken = true
					break corrupted
				} else {
					stack, _ = pop(stack)
				}
			} else {
				stack = push(stack, c)
			}
		}
		if !broken {
			incompleteLines = append(incompleteLines, line)
		}
	}
	fmt.Println("Part 1:", totalScore)
	return incompleteLines
}

func part2(lines []string) {
	scoreFor := map[rune]int64{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	var lineScores []int64
	var stack []rune
	for _, line := range lines {
		for _, c := range line {
			if len(stack) == 0 {
				stack = push(stack, c)
			} else if isClosingBrace[c] {
				stack, _ = pop(stack)
			} else {
				stack = push(stack, c)
			}
		}
		var lineScore int64
		for len(stack) > 0 {
			var c rune
			stack, c = pop(stack)
			lineScore = lineScore*5 + scoreFor[c]
		}
		lineScores = append(lineScores, lineScore)
	}
	sort.Slice(lineScores, func(i, j int) bool {
		return lineScores[i] < lineScores[j]
	})
	fmt.Println("Part 2:", lineScores[len(lineScores)/2])
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("bruh")
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	incompleteLines := part1(lines)
	part2(incompleteLines)
}
