package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/roessland/gopkg/mathutil"
)

func main() {
	input := ReadInput()
	network := MakeNetwork(input.Nodes)

	part1(input.Instructions, network)
	part2(input.Instructions, network, 99)
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

type State struct {
	Instrs  *string
	Network map[string]Node
	Locs    []string
	Steps   int
}

func (s State) IsTarget() bool {
	for _, loc := range s.Locs {
		if !strings.HasSuffix(loc, "Z") {
			return false
		}
	}
	return true
}

func (s State) Ip() int {
	return s.Steps % len(*s.Instrs)
}

func (s State) Instr() byte {
	return (*s.Instrs)[s.Ip()]
}

func (s State) ID() string {
	return strings.Join(s.Locs, ",") + fmt.Sprintf(":%d", s.Ip())
}

func (s State) Substate(locIdx int) State {
	return State{
		Instrs:  s.Instrs,
		Network: s.Network,
		Locs:    []string{s.Locs[locIdx]},
		Steps:   s.Steps,
	}
}

func DetectCycle(initialState State) (state State, cycleLen int) {
	visited := map[string]bool{}
	lastVisited := map[string]int{}
	state = initialState
	for {
		if visited[state.ID()] {
			return state, state.Steps - lastVisited[state.ID()]
		}
		visited[state.ID()] = true
		lastVisited[state.ID()] = state.Steps
		state = state.Next()
	}
}

func (state State) Next() State {
	nextLocs := make([]string, len(state.Locs))
	for i, loc := range state.Locs {
		if state.Instr() == 'L' {
			nextLocs[i] = state.Network[loc].Left
		} else {
			nextLocs[i] = state.Network[loc].Right
		}
	}
	return State{
		Locs:    nextLocs,
		Steps:   state.Steps + 1,
		Instrs:  state.Instrs,
		Network: state.Network,
	}
}

// FindCycleTargets find the number of steps from a position in the cycle
// to all Zs in the cycle. Incidentally there is always 1.
func (state State) FindCycleTargets() []int {
	targetOffsets := []int{}
	initialSteps := state.Steps
	initialStateID := state.ID()
	state = state.Next()
	for state.ID() != initialStateID {
		if strings.HasSuffix(state.Locs[0], "Z") {
			targetOffsets = append(targetOffsets, state.Steps-initialSteps)
		}
		state = state.Next()
	}
	return targetOffsets
}

func part2(instructions string, n map[string]Node, maxLocs int) {
	locs := []string{}
	for _, node := range n {
		if strings.HasSuffix(node.ID, "A") {
			locs = append(locs, node.ID)
		}
	}
	sort.Strings(locs)
	if len(locs) > maxLocs {
		locs = locs[:maxLocs]
	}

	globalState := State{
		Locs:    locs,
		Steps:   0,
		Instrs:  &instructions,
		Network: n,
	}

	fmt.Println("Solving with", len(locs), "locations")

	fmt.Println("Detecting cycles for each start")
	cycleLocs := []State{}
	cycleLens := []int{}
	for i := range locs {
		substate := globalState.Substate(i)
		cycleEntryState, cycleLen := DetectCycle(substate)
		cycleLocs = append(cycleLocs, cycleEntryState)
		cycleLens = append(cycleLens, cycleLen)
		fmt.Println("  ", cycleEntryState.Locs, "cycle start at", cycleEntryState.Steps)
		fmt.Println("  ", cycleLen, "cycle len")
	}

	fmt.Println("Aligning cycle entries")
	maxCycleEntrySteps := 0
	for _, cycleEntryState := range cycleLocs {
		if cycleEntryState.Steps > maxCycleEntrySteps {
			maxCycleEntrySteps = cycleEntryState.Steps
		}
	}
	for i := range cycleLocs {
		for cycleLocs[i].Steps < maxCycleEntrySteps {
			cycleLocs[i] = cycleLocs[i].Next()
		}
		fmt.Println(cycleLocs[i].Steps, "cycle entry at")
	}

	fmt.Println("Verifying each cycle only has one Z")
	cycleTargets := []int{}
	for i := range cycleLocs {
		cycleTargets_ := cycleLocs[i].FindCycleTargets()
		if len(cycleTargets_) != 1 {
			panic("expected one cycle target")
		}
		cycleTargets = append(cycleTargets, cycleTargets_[0])
	}

	fmt.Println("Verifying cycle targets")
	for i := range cycleLocs {
		s := cycleLocs[i]
		for j := 0; j < cycleTargets[i]; j++ {
			s = s.Next()
		}
		if !s.IsTarget() {
			panic("expected target")
		}
		// fmt.Println("Steps from cycle start", s.Steps-cycleLocs[i].Steps)
	}

	// steps - cycleLoc[i].Steps = cycleTargets[i] (mod cycleLens[i])
	// steps = cycleTargets[i] + cycleLoc[i].Steps (mod cycleLens[i])
	fmt.Println("Building CRT system")
	as := []int64{}
	ns := []int64{}
	for i := range cycleLocs {
		a := (cycleLocs[i].Steps + cycleTargets[i]) % cycleLens[i]
		as = append(as, int64(a))
		ns = append(ns, int64(cycleLens[i]))
		fmt.Printf("  s = %d mod %d\n", a, cycleLens[i])
	}

	fmt.Println("Simplifying CRT system")
	as, ns = simplifyCRT(as, ns)
	for i := range as {
		fmt.Printf("  s = %d mod %d\n", as[i], ns[i])
	}

	fmt.Println("Solving CRT system")
	s, N := crt(as, ns)
	for s < int64(maxCycleEntrySteps) {
		s += N
	}
	fmt.Println("Part 2 (Chinese remainder theorem):", s)

	// fmt.Println("Brute forcing part 2 (takes forever)")
	// for !globalState.IsTarget() {
	// 	globalState = globalState.Next()
	// }
	//
	// fmt.Println("Part 2 (Bruteforce):", globalState.Steps)
}

func simplifyCRT(as, ns []int64) ([]int64, []int64) {
	// Make input numbers coprime with each other.
	as_ := make([]int64, len(as))
	ns_ := make([]int64, len(ns))
	copy(as_, as)
	copy(ns_, ns)

	// If any coprime pairs exist, divide one of them by their GCD,
	// until no coprime pairs remain.
outer:
	for {
		for i := range ns_ {
			for j := range ns_ {
				if i == j {
					continue
				}
				gcd := mathutil.GCD(ns_[i], ns_[j])
				if gcd != 1 {
					as_[i] /= gcd
					ns_[i] /= gcd
					continue outer
				}
			}
		}
		break
	}
	return as_, ns_
}

func assertPairwiseCoprime(ns []int64) {
	for i := range ns {
		for j := range ns {
			if i == j {
				continue
			}
			gcd := mathutil.GCD(ns[i], ns[j])
			if gcd != 1 {
				panic("ns are not coprime")
			}
		}
	}
}

// crt solves the system of congruences using the chinese remainder theorem.
//
// x = a1 (mod n1)
// x = a2 (mod n2)
// ...
// x = ak (mod nk)
//
// N = prod(n1 ... nk)
//
// https://brilliant.org/wiki/chinese-remainder-theorem/
//
// The solution x is unique mod N.
func crt(as []int64, ns []int64) (x, N int64) {
	// Validate as_ are less than ns_.
	for i := range as {
		if as[i] >= ns[i] {
			panic("as_ must be less than ns_")
		}
	}

	assertPairwiseCoprime(ns)

	N = int64(1)
	for _, n := range ns {
		N *= n
	}

	ys := make([]int64, len(as))
	for i := range ns {
		ys[i] = N / ns[i]
	}

	zs := make([]int64, len(as))
	for i := range ns {
		zs[i] = mathutil.ModularInverse(ys[i], ns[i])
	}

	x = int64(0)
	for i := range ns {
		x = (x + as[i]*ys[i]*zs[i]) % N
	}

	return x, N
}

func detectCycleTargets(instrs string, ip0 int, cycleStart string, n map[string]Node) (int, int) {
	cycleSteps := 0
	ip0 = ip0 % len(instrs)
	loc := cycleStart
	dsts := []int{}

	ip := ip0
	for cycleSteps == 0 || loc != cycleStart || ip != ip0 {
		if strings.HasSuffix(loc, "Z") {
			dsts = append(dsts, cycleSteps)
		}
		if instrs[ip] == 'L' {
			loc = n[loc].Left
		} else {
			loc = n[loc].Right
		}
		ip = (ip + 1) % len(instrs)
		cycleSteps += 1
	}
	if len(dsts) != 1 {
		panic("expected one destination")
	}
	return cycleSteps, dsts[0]
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
