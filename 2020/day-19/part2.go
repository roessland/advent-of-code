package main

/*
Wow, that was a difficult one for me.
I got stuck on trying to return only the longest match,
and had `type Rule func(s string)int` instead of returning set.
 */

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule func(s string, i int) (sizes Set)

type Rules map[string]Rule

type Set map[int]bool

func SetUnion(a, b Set) Set {
	u := make(Set)
	for k, v := range a {
		if !v {
			continue
		}
		u[k] = v
	}
	for k, v := range b {
		if !v {
			continue
		}
		u[k] = v
	}
	return u
}

func SetContains(a Set, i int) bool {
	for k := range a {
		if k == i {
			return true
		}
	}
	return false
}

func SetEqual(a, b Set) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !b[i] {
			return false
		}
	}
	for i := range b {
		if !a[i] {
			return false
		}
	}
	return true
}

func Const(c uint8) Rule {
	return func(s string, i int) Set {
		if i < len(s) && s[i] == c {
			return Set{1: true}
		}
		return nil
	}
}

func Ref(rules Rules, ruleName string) Rule {
	return func(s string, i int) Set {
		return rules[ruleName](s, i)
	}
}

func Seq(rs []Rule) Rule {
	if len(rs) == 1 {
		return Seq1(rs[0])
	} else if len(rs) == 2 {
		return Seq2(rs[0], rs[1])
	} else if len(rs) == 3 {
		return Seq3(rs[0], rs[1], rs[2])
	} else {
		panic("unsupported seq len")
	}
}

func Seq1(r Rule) Rule {
	return r
}

func Seq2(rA, rB Rule) Rule {
	return func(s string, i int) Set {
		sizesA := rA(s, i)
		if len(sizesA) == 0 {
			return nil
		}
		sizes := make(Set)
		for szA := range sizesA {
			sizesB := rB(s, i+szA)
			for szB := range sizesB {
				sizes[szA+szB] = true
			}
		}
		return sizes
	}
}

func Seq3(rA, rB, rC Rule) Rule {
	return func(s string, i int) Set {
		sizesA := rA(s, i)
		if len(sizesA) == 0 {
			return nil
		}
		sizes := make(Set)
		for szA := range sizesA {
			sizesB := rB(s, i+szA)
			if len(sizesB) == 0 {
				return nil
			}
			for szB := range sizesB {
				sizesC := rC(s, i+szA+szB)
				if len(sizesC) == 0 {
					return nil
				}
				for szC := range sizesC {
					sizes[szA+szB+szC] = true
				}
			}
		}
		return sizes
	}
}

func Or(rs []Rule) Rule {
	return func(s string, i int) Set {
		allSizes := make(Set)
		for _, r := range rs {
			sizes := r(s, i)
			allSizes = SetUnion(allSizes, sizes)
		}
		return allSizes
	}
}

func (rules Rules) AddRule(str string) {
	as := strings.Split(str, ": ")
	ruleName := as[0]

	if strings.HasPrefix(as[1], `"`) {
		c := as[1][1]
		rules[ruleName] = Const(c)
		return
	}

	var orRules []Rule
	for _, rule := range strings.Split(as[1], " | ") {
		refRuleNames := strings.Split(rule, " ")
		var seqRules []Rule
		for _, refRuleName := range refRuleNames {
			seqRules = append(seqRules, Ref(rules, refRuleName))
		}
		orRules = append(orRules, Seq(seqRules))
	}
	rules[ruleName] = Or(orRules)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// Read rules
	rules := make(Rules)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if line == "8: 42" {
			line = "8: 42 | 42 8"
		}
		if line == "11: 42 31" {
			line = "11: 42 31 | 42 11 31"
		}
		rules.AddRule(line)
	}

	// Read messages
	matches := 0
	for scanner.Scan() {
		msg := scanner.Text()
		sizes := rules["0"](msg, 0)
		if SetContains(sizes, len(msg)) {
			matches++
		}
	}
	fmt.Println("Part 2:", matches)
}
