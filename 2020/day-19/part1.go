package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Rule func(s string) int


type Rules map[string]Rule

func (rules Rules) AddRule(str string) {
	as := strings.Split(str, ": ")
	ruleName := as[0]
	if strings.HasPrefix(as[1], `"`) {
		// Static rule
		c := as[1][1]
		rules[ruleName] = func(s string)int{
			if len(s) > 0 && s[0] == c {
				return 1
			}
			return 0
		}
	} else {
		// Subrules, including OR
		var orRules []Rule
		for _, rule := range strings.Split(as[1], " | ") {
			subRuleNames := strings.Split(rule, " ")
			orRules = append(orRules, func(s string)int {
				j := 0
				for _, subRuleName := range subRuleNames {
					i := rules[subRuleName](s)
					if i == 0 {
						return 0
					}
					j += i
					s = s[i:]
				}
				return j
			})
		}
		rules[ruleName] = func(s string)int{
			for _, orRule := range orRules {
				i := orRule(s)
				if i > 0 {
					return i
				}
			}
			return 0
		}
	}
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
		rules.AddRule(line)
	}

	// Read messages
	matches := 0
	for scanner.Scan() {
		msg := scanner.Text()
		if rules["0"](msg) == len(msg) {
			matches++
		}
	}
	fmt.Println(matches)
}
