package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	input := ReadInput()
	part1(input)
}

type PartRating map[string]int

type Workflow struct {
	In    chan PartRating
	Rules []Rule
}

func (wf *Workflow) Run() {
nextPart:
	for part := range wf.In {
		for _, rule := range wf.Rules {
			if next := rule(part); next != nil {
				next <- part
				continue nextPart
			}
		}
		e := fmt.Sprintf("no rule matched for part %v", part, wf.Rules)
		panic(e)
	}
}

type Rule func(PartRating) chan<- PartRating

func Op(op string, left, right int) bool {
	if op == "<" {
		return left < right
	}
	if op == ">" {
		return left > right
	}
	panic("unknown operator")
}

func part1(input Input) {
	chans := make(map[string]chan PartRating)
	workflows := make(map[string]*Workflow)

	for _, workflowInput := range input.Workflows {
		rules := []Rule{}
		for _, ruleInput := range workflowInput.Rules {
			ruleInput := ruleInput
			rules = append(rules, func(part PartRating) chan<- PartRating {
				if ruleInput.Operator == "" || Op(ruleInput.Operator, part[ruleInput.LeftCategory], ruleInput.RightThreshold) {
					return chans[ruleInput.Output]
				}
				return nil
			})
		}
		ch := make(chan PartRating)
		chans[workflowInput.WorkflowName] = ch
		workflows[workflowInput.WorkflowName] = &Workflow{Rules: rules, In: ch}
	}

	for _, wf := range workflows {
		go wf.Run()
	}

	chans["R"] = make(chan PartRating)
	chans["A"] = make(chan PartRating)

	go func() {
		for _, part := range input.Parts {
			chans["in"] <- part
		}
	}()

	done := 0
	acceptedParts := []PartRating{}
	for done < len(input.Parts) {
		select {
		case part := <-chans["R"]:
			fmt.Println("rejected", part)
		case part := <-chans["A"]:
			fmt.Println("accepted", part)
			acceptedParts = append(acceptedParts, part)
		}
		done++
	}

	for _, ch := range chans {
		close(ch)
	}

	sum := 0
	for _, part := range acceptedParts {
		sum += part["x"] + part["m"] + part["a"] + part["s"]
	}
	fmt.Println(sum)
}

type Input struct {
	Workflows []WorkflowInput
	Parts     []PartRating
}

type WorkflowInput struct {
	WorkflowName string
	Rules        []RuleInput
}

type RuleInput struct {
	Output         string
	LeftCategory   string
	Operator       string
	RightThreshold int
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
					LeftCategory:   ruleMatches[1],
					Operator:       ruleMatches[2],
					RightThreshold: aocutil.Atoi(ruleMatches[3]),
					Output:         ruleMatches[4],
				})
			}
		}
		workflows = append(workflows, WorkflowInput{
			WorkflowName: workflowName,
			Rules:        rules,
		})
	}

	parts := []PartRating{}
	partRe := regexp.MustCompile(`(\w)=(\d+)`)
	for ; i < len(lines); i++ {
		partRatingStr := strings.Trim(lines[i], "{}")
		part := make(PartRating)
		for _, partRating := range strings.Split(partRatingStr, ",") {
			matches := partRe.FindStringSubmatch(partRating)
			part[matches[1]] = aocutil.Atoi(matches[2])
		}
		parts = append(parts, part)
	}

	return Input{
		Workflows: workflows,
		Parts:     parts,
	}
}
