package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ticket []int

func NewTicket(line string) Ticket {
	t := Ticket{}
	numsStr := strings.Split(line, ",")
	for _, numStr := range numsStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		t = append(t, num)
	}
	return t
}

type Rule struct {
	Field  string
	Ranges []Range
}

func (r Rule) Valid(n int) bool {
	for _, rang := range r.Ranges {
		if rang.Contains(n) {
			return true
		}
	}
	return false
}

type Range struct {
	Lo, Hi int
}

func NewRange(rule string) Range {
	loHi := strings.Split(rule, "-")
	lo, err1 := strconv.Atoi(loHi[0])
	hi, err2 := strconv.Atoi(loHi[1])
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
	return Range{lo, hi}
}

func (r Range) Contains(n int) bool {
	return r.Lo <= n && n <= r.Hi
}

func NewRule(line string) Rule {
	var r Rule
	depRules := strings.Split(line, ": ")
	r.Field = depRules[0]
	ranges := strings.Split(depRules[1], " or ")
	for _, rangeStr := range ranges {
		r.Ranges = append(r.Ranges, NewRange(rangeStr))
	}
	return r
}

type Input struct {
	Rules   []Rule
	Your    Ticket
	Nearbys []Ticket
}

func ReadInput() Input {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input Input
	scanner := bufio.NewScanner(f)
	stage := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			stage++
			continue
		}
		switch stage {
		case 0:
			input.Rules = append(input.Rules, NewRule(line))
		case 1:
			input.Your = NewTicket(line)
		case 2:
			input.Nearbys = append(input.Nearbys, NewTicket(line))
		}
	}
	return input
}

func (input Input) ValidForSomeRule(n int) bool {
	for _, rule := range input.Rules {
		if rule.Valid(n) {
			return true
		}
	}
	return false
}

func (input Input) Part1() []Ticket {
	var validTickets []Ticket
	errorRate := 0
	for _, t := range input.Nearbys {
		hasError := false
		for _, n := range t {
			if !input.ValidForSomeRule(n) {
				hasError = true
				errorRate += n
			}
		}
		if !hasError {
			validTickets = append(validTickets, t)
		}
	}
	fmt.Println("Part 1:", errorRate)
	return validTickets
}

func clone(m map[int]struct{}) map[int]struct{} {
	c := map[int]struct{}{}
	for key := range m {
		c[key] = struct{}{}
	}
	return c
}

func (in Input) Search(remainingRuleIdxs map[int]struct{}, remainingTicketCols map[int]struct{}) (success bool, validRuleIdxs []int, validTicketCols []int) {
	// We did it!
	if len(remainingRuleIdxs) == 0 {
		return true, []int{}, []int{}
	}

	// Group rules by which columns they are valid for
	validColsForRule := make(map[int][]int)
	for ruleIdx := range remainingRuleIdxs {
		for col := range remainingTicketCols {
			validForAllTickets := true
			for _, ticket := range in.Nearbys {
				if !in.Rules[ruleIdx].Valid(ticket[col]) {
					validForAllTickets = false
					break
				}
			}
			if validForAllTickets {
				validColsForRule[ruleIdx] = append(validColsForRule[ruleIdx], col)
			}
		}
	}

	// Find the column with the minimum number of valid rules
	minColsRuleIdx := -1
	minCols := math.MaxInt32
	for ruleIdx, validCols := range validColsForRule {
		numValidCols := len(validCols)
		if numValidCols == 0 {
			return false, nil, nil // descended into unsolvable branch
		}
		if numValidCols < minCols {
			minCols = numValidCols
			minColsRuleIdx = ruleIdx
		}
	}

	// Recurse for each possible column for that rule. One of them should be the correct choice.
	rules := clone(remainingRuleIdxs)
	delete(rules, minColsRuleIdx)
	for _, col := range validColsForRule[minColsRuleIdx] {
		cols := clone(remainingTicketCols)
		delete(cols, col)
		valid, subRules, subCols := in.Search(rules, cols)
		if valid {
			return valid, append(subRules, minColsRuleIdx), append(subCols, col)
		}
	}

	return false, nil, nil
}

func (input Input) Part2() {
	remainingRuleIdxsSet := map[int]struct{}{}
	remainingTicketColsSet := map[int]struct{}{}
	for i, _ := range input.Rules {
		remainingRuleIdxsSet[i] = struct{}{}
		remainingTicketColsSet[i] = struct{}{}
	}

	_, ruleIdxs, colIdxs := input.Search(remainingRuleIdxsSet, remainingTicketColsSet)

	// Four your own ticket, find fields starting with "departure" and multiply them.
	prod := 1
	for i := 0; i < len(ruleIdxs); i++ {
		rule := input.Rules[ruleIdxs[i]]
		col := colIdxs[i]
		if !rule.Valid(input.Your[col]) {
			panic("fuck")
		}
		if strings.HasPrefix(rule.Field, "departure") {
			prod *= input.Your[col]
		}
	}
	fmt.Println("Part 2:", prod)
}

func main() {
	input := ReadInput()
	input.Nearbys = input.Part1()
	input.Part2()
}
