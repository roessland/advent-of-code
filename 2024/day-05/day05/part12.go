package day05

import (
	"embed"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Rule struct {
	Fst int
	Snd int
}

type Update = []int

type Graph map[int]map[int]bool

func ReadInput(inputName string) ([]Rule, []Update) {
	var rules []Rule
	var updates []Update
	lines := aocutil.ReadFileAsInts(Input, inputName)
	for _, line := range lines {
		if len(line) == 2 {
			rules = append(rules, Rule{line[0], line[1]})
		} else if len(line) == 0 {
			continue
		} else {
			updates = append(updates, line)
		}
	}
	return rules, updates
}

// HasCycle returns true if the graph contains a cycle.
func HasCycle(g Graph) (hasCycle bool) {
	defer func() {
		if r := recover(); r != nil {
			hasCycle = true
		}
	}()
	TopoSort(g)
	return false
}

// Merge returns a graph containing all edges and nodes from both input graphs.
func Merge(gA Graph, gB Graph) Graph {
	// Keep edges in update graph
	relevantGraph := make(Graph)
	for i := range gA {
		if _, ok := relevantGraph[i]; !ok {
			relevantGraph[i] = make(map[int]bool)
		}
		for j := range gA[i] {
			relevantGraph[i][j] = true
		}
	}
	for i := range gB {
		if _, ok := relevantGraph[i]; !ok {
			relevantGraph[i] = make(map[int]bool)
		}
		for j := range gB[i] {
			relevantGraph[i][j] = true
		}
	}
	return relevantGraph
}

// TopoSort returns a topological sorting of the graph nodes, or panics if there is a cycle.
func TopoSort(g Graph) []int {
	permMarked := make(map[int]bool)
	tempMarked := make(map[int]bool)
	L := make([]int, 0)

	var visit func(int)
	visit = func(n int) {
		if permMarked[n] {
			return
		}
		if tempMarked[n] {
			panic("cycle")
		}

		tempMarked[n] = true
		for m := range g[n] {
			visit(m)
		}

		permMarked[n] = true
		L = append(L, n)
	}

	selectUnmarked := func() int {
		for i := range g {
			if !permMarked[i] && !tempMarked[i] {
				return i
			}
		}
		panic("No unmarked nodes")
	}

	for len(permMarked) < len(g) {
		n := selectUnmarked()
		visit(n)
	}
	return L
}

// RulesGraph creates a graph representing all the rules
func RulesGraph(rules []Rule) Graph {
	g := make(Graph)
	for _, rule := range rules {
		if _, ok := g[rule.Fst]; !ok {
			g[rule.Fst] = make(map[int]bool)
		}
		g[rule.Fst][rule.Snd] = true
	}
	return g
}

// UpdateGraph creates a graph representing an update
func UpdateGraph(update Update) Graph {
	g := make(Graph)
	for i := 0; i < len(update)-1; i++ {
		if _, ok := g[update[i]]; !ok {
			g[update[i]] = make(map[int]bool)
		}
		if _, ok := g[update[i+1]]; !ok {
			g[update[i+1]] = make(map[int]bool)
		}
		g[update[i]][update[i+1]] = true
	}
	return g
}

// PickNodes return a subgraph with only the given nodes
func PickNodes(g Graph, nodes []int) Graph {
	gSub := make(Graph)
	for _, i := range nodes {
		if _, ok := gSub[i]; !ok {
			gSub[i] = make(map[int]bool)
		}
		for _, j := range nodes {
			if g[i][j] {
				gSub[i][j] = true
			}
		}
	}
	return gSub
}

func Part12(inputName string) (int, int) {
	rules, updates := ReadInput(inputName)
	rulesGraph := RulesGraph(rules)

	sum1 := 0
	sum2 := 0
	for _, update := range updates {
		updateGraph := UpdateGraph(update)
		relevantRulesGraph := PickNodes(rulesGraph, update)

		if HasCycle(Merge(updateGraph, relevantRulesGraph)) {
			sortedUpdate := TopoSort(relevantRulesGraph)
			sum2 += sortedUpdate[len(sortedUpdate)/2]
		} else {
			sum1 += update[len(update)/2]
		}
	}

	return sum1, sum2
}
