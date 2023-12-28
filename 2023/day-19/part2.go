package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

const (
	MinRating = 1
	MaxRating = 4000
)

func main() {
	input := ReadInput()
	part2(input)
}

type Set map[int]bool

type Sets map[string]Set

func NewSet() Set {
	s := make(Set)
	for i := 1; i <= MaxRating; i++ {
		s[i] = true
	}
	return s
}

func (s Set) Union(other Set) Set {
	union := make(Set)
	for k := range s {
		union[k] = true
	}
	for k := range other {
		union[k] = true
	}
	return union
}

func (s Set) Intersect(other Set) Set {
	intersection := make(Set)
	for k := range s {
		if other[k] {
			intersection[k] = true
		}
	}
	return intersection
}

func (s Set) Difference(other Set) Set {
	difference := make(Set)
	for k := range s {
		if !other[k] {
			difference[k] = true
		}
	}
	return difference
}

func NewSets() Sets {
	return Sets{
		"x": NewSet(),
		"m": NewSet(),
		"a": NewSet(),
		"s": NewSet(),
	}
}

func (sts Sets) ShallowCopy() Sets {
	newSets := make(Sets)
	for cat, set := range sts {
		newSets[cat] = set
	}
	return newSets
}

func (pr Sets) Cardinality() int {
	prod := 1
	for _, category := range pr {
		prod *= len(category)
	}
	return prod
}

type Workflow struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	Set
	OutName string
	Cat     string
}

func Op(op string, left, right int) bool {
	if op == "<" {
		return left < right
	}
	if op == ">" {
		return left > right
	}
	panic(fmt.Sprintf("Unknown op: %v", op))
}

func MakeRule(ruleInput RuleInput) Rule {
	set := make(Set)
	for included := range NewSet() {
		if ruleInput.Op == "" || Op(ruleInput.Op, included, ruleInput.Threshold) {
			set[included] = true
		}
	}
	return Rule{
		Set:     set,
		OutName: ruleInput.Output,
		Cat:     ruleInput.Cat,
	}
}

func N(wfs map[string]Workflow, wfName string, sets Sets) int {
	if wfName == "A" {
		return sets.Cardinality()
	}

	if wfName == "R" {
		return 0
	}

	sum := 0
	wf := wfs[wfName]
	for _, rule := range wf.Rules {
		if rule.Cat == "" {
			sum += N(wfs, rule.OutName, sets)
			return sum
		}
		cat := rule.Cat
		set := sets[cat]

		left := set.Intersect(rule.Set)
		leftSets := sets.ShallowCopy()
		leftSets[cat] = left
		sum += N(wfs, rule.OutName, leftSets)

		right := set.Difference(rule.Set)
		rightSets := sets.ShallowCopy()
		rightSets[cat] = right
		sets = rightSets
	}
	if sets.Cardinality() > 0 {
		panic(fmt.Sprintf("Sets not empty: %v", sets))
	}
	return sum
}

func part2(input Input) {
	// Make workflows
	workflows := map[string]Workflow{}
	for _, workflowInput := range input.Workflows {
		rules := []Rule{}
		for _, ruleInput := range workflowInput.Rules {
			rules = append(rules, MakeRule(ruleInput))
		}

		workflows[workflowInput.WorkflowName] = Workflow{
			Name:  workflowInput.WorkflowName,
			Rules: rules,
		}
	}

	pr := map[string]Set{
		"x": NewSet(),
		"m": NewSet(),
		"a": NewSet(),
		"s": NewSet(),
	}

	fmt.Println("Part 2: ", N(workflows, "in", pr))
}

type Input struct {
	Workflows []WorkflowInput
	Parts     []Sets
}

type WorkflowInput struct {
	WorkflowName string
	Rules        []RuleInput
}

type RuleInput struct {
	Output    string
	Cat       string
	Op        string
	Threshold int
}

func ReadInput() Input {
	lines := aocutil.ReadLines("input.txt")

	workflows := []WorkflowInput{}
	workflowNameRe := regexp.MustCompile(`^(\w+){(.*)}`)
	ruleRe := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)
	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}
		workflowStr := lines[i]
		nameAndRulesMatches := workflowNameRe.FindStringSubmatch(workflowStr)
		workflowName, ruleStrs := nameAndRulesMatches[1], nameAndRulesMatches[2]
		rules := []RuleInput{}
		for _, ruleStr := range strings.Split(ruleStrs, ",") {
			ruleMatches := ruleRe.FindStringSubmatch(ruleStr)
			if len(ruleMatches) == 0 {
				rules = append(rules, RuleInput{Output: ruleStr})
			} else {
				rules = append(rules, RuleInput{
					Cat:       ruleMatches[1],
					Op:        ruleMatches[2],
					Threshold: aocutil.Atoi(ruleMatches[3]),
					Output:    ruleMatches[4],
				})
			}
		}
		workflows = append(workflows, WorkflowInput{
			WorkflowName: workflowName,
			Rules:        rules,
		})
	}

	return Input{
		Workflows: workflows,
	}
}
