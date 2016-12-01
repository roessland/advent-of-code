package main

import "io/ioutil"
import "fmt"
import "strings"
import "strconv"

type Properties map[string]int
type Requirements map[string]int

func (p Properties) Match1(r Requirements) bool {
	for rkey, rval := range r {
		if pval, ok := p[rkey]; ok && pval != rval {
			return false
		}
	}
	return true
}

func (p Properties) Match2(r Requirements) bool {
	for rkey, rval := range r {
		if pval, ok := p[rkey]; ok {
			switch rkey {
			case "cats":
				fallthrough
			case "trees":
				if pval <= rval {
					return false
				}
			case "pomeranians":
				fallthrough
			case "goldfish":
				if pval >= rval {
					return false
				}
			default:
				if pval != rval {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	requirements := Requirements{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	buf, _ := ioutil.ReadFile("input.txt")
	for i, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		properties := Properties{}
		for _, keyval := range strings.Split(line[strings.Index(line, ": ")+2:], ", ") {
			colon := strings.Index(keyval, ": ")
			key := keyval[:colon]
			val, _ := strconv.Atoi(keyval[colon+2:])
			properties[key] = val
		}
		if properties.Match1(requirements) {
			fmt.Printf("Part 1: Sue %v is a match\n", i+1)
		}
		if properties.Match2(requirements) {
			fmt.Printf("Part 2: Sue %v is a match\n", i+1)
		}
	}
}
