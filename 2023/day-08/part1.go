package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func main() {
	input := ReadInput()
	network := MakeNetwork(input.Nodes)

	// part1(input.Instructions, network)
	part2(input.Instructions, network)
}

func MakeNetwork(nodes []NodeInput) map[string]Node {
	n := map[string]Node{}
	for _, node := range nodes {
		n[node.ID] = Node(node)
	}
	return n
}

type Node struct {
	ID    string
	Left  string
	Right string
}

func part1(instructions string, network map[string]Node) {
	loc := "AAA"
	ip := 0
	steps := 0
	for loc != "ZZZ" {
		if instructions[ip] == 'L' {
			loc = network[loc].Left
		} else {
			loc = network[loc].Right
		}

		steps += 1
		ip = (ip + 1) % len(instructions)
	}
	fmt.Println(steps)
}

func part2(instructions string, n map[string]Node) {
	locs := []string{}
	for _, node := range n {
		if strings.HasSuffix(node.ID, "A") {
			locs = append(locs, node.ID)
		}
	}
	fmt.Println("starting at", locs)

	fmt.Println("detecting cycle")

	for _, loc := range locs {
		stepsUntilCycleStart, cycleStart := detectCycle(instructions, n, loc)
		cycleLen, _ := detectCycleTargets(instructions, stepsUntilCycleStart, cycleStart, n)
		// loc(stepsUntilCycleStart + cycleDst) = Z
		// loc(stepsUntilCycleStart + (cycleDst + cycleLen * x) % cycleLen) = Z
		// Let x be total steps since start: x = stepsUntilCycleStart + cycleDst + cycleLen * n
		// locA(x) = Z
		// locB(x) = Z
		// x = stepsUntilCycleStartA + cycleDstA + cycleLenA * n = stepsUntilCycleStartA (mod cycleLen A)
		// cycleLenA * n = - cycleDstA (mod cycleLenA)
		// cycleLenB * n = - cycleDstB (mod cycleLenB)
		//
		// n = - cycleDstA  (mod cycleLenA)
		// n = - cycleDstB  (mod cycleLenB)
		// x - stepsUntilCycleStart = 0 (mod cycleLenA)
		fmt.Printf("x = %d (mod %d)\n", stepsUntilCycleStart%cycleLen, cycleLen)
		// Is Z when s % cycleLen = 0
	}

	ip := 0
	steps := 0

	for !allLocsEndWithZ(locs) {
		if (steps-583)%1241 == 0 {
			locss := []string{}
			locss = append(locss, locs...)
			sort.Strings(locss)
			fmt.Println(locss)

			fmt.Println(locs)
		}
		for i, loc := range locs {
			if instructions[ip] == 'L' {
				locs[i] = n[loc].Left
			} else {
				locs[i] = n[loc].Right
			}
		}
		steps += 1
		ip = (ip + 1) % len(instructions)
		if steps%100000 == 0 {
			fmt.Println(steps)
		}
	}

	fmt.Println(steps)
}

func detectCycleTargets(instrs string, ip int, cycleStart string, n map[string]Node) (int, int) {
	cycleSteps := 0
	ip = ip % len(instrs)
	loc := cycleStart
	dsts := []int{}

	cycle := []string{}
	for cycleSteps == 0 || loc != cycleStart {
		cycle = append(cycle, loc)
		if instrs[ip] == 'L' {
			loc = n[loc].Left
		} else {
			loc = n[loc].Right
		}
		if strings.HasSuffix(loc, "Z") {
			dsts = append(dsts, cycleSteps)
		}
		ip = (ip + 1) % len(instrs)
		cycleSteps += 1
	}
	if len(dsts) != 1 {
		panic("expected one destination")
	}
	return cycleSteps, dsts[0]
}

func detectCycle(instrs string, n map[string]Node, start string) (int, string) {
	turtle := start
	var hare string
	if instrs == "L" {
		hare = n[start].Left
	} else {
		hare = n[start].Right
	}

	ipTurtle := 0
	ipHare := 1
	stepsTurtle := 0
	for turtle != hare || ipTurtle != ipHare {
		if instrs[ipTurtle] == 'L' {
			turtle = n[turtle].Left
		} else {
			turtle = n[turtle].Right
		}
		stepsTurtle += 1
		ipTurtle = (ipTurtle + 1) % len(instrs)

		if instrs[ipHare] == 'L' {
			hare = n[hare].Left
		} else {
			hare = n[hare].Right
		}
		ipHare = (ipHare + 1) % len(instrs)

		if instrs[ipHare] == 'L' {
			hare = n[hare].Left
		} else {
			hare = n[hare].Right
		}
		ipHare = (ipHare + 1) % len(instrs)
	}
	return stepsTurtle, hare
}

func allLocsEndWithZ(locs []string) bool {
	for _, loc := range locs {
		if !strings.HasSuffix(loc, "Z") {
			return false
		}
	}
	return true
}

type Input struct {
	Instructions string
	Nodes        []NodeInput
}

type NodeInput struct {
	ID    string
	Left  string
	Right string
}

func ReadInput() Input {
	input := Input{}

	lines := aocutil.ReadLines("input.txt")

	input.Instructions = lines[0]
	for _, line := range lines[2:] {
		lineRegex := regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)
		matches := lineRegex.FindStringSubmatch(line)
		input.Nodes = append(input.Nodes, NodeInput{
			ID:    matches[1],
			Left:  matches[2],
			Right: matches[3],
		})
	}
	return input
}

// 142422532599150720  too high
// 472*N = -15516 (mod M)
// 366*N = -16042 (mod M)
// 395*N = -20776 (mod M)
// 2556*N = -18672 (mod M)
// 658*N = -12360 (mod M)
// 1241*N = -19198 (mod M)
// M = 142422532599150720
