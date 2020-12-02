package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Min, Max int
	Letter rune
}

func NewRuleFromString(ruleStr string) Rule {
	intervalLetter := strings.Split(ruleStr, " ")
	intervalStr := intervalLetter[0]
	letter := rune(intervalLetter[1][0])
	minMaxStr := strings.Split(intervalStr, "-")
	min, err1 := strconv.Atoi(minMaxStr[0])
	max, err2 := strconv.Atoi(minMaxStr[1])
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
	return Rule{Min: min, Max: max, Letter: letter}
}

func (rule Rule) Valid(s string) bool {
	letterCount := 0
	for _, r := range s {
		if r == rule.Letter {
			letterCount++
		}
	}
	return rule.Min <= letterCount && letterCount <= rule.Max
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	numValid := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rulePassword := strings.Split(line, ": ")
		ruleStr := rulePassword[0]
		password := rulePassword[1]
		rule := NewRuleFromString(ruleStr)
		if rule.Valid(password) {
			numValid++
		}
	}
	fmt.Println(numValid)
}