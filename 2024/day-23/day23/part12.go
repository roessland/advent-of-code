package day23

import (
	"embed"
	"sort"
	"strings"

	"github.com/roessland/advent-of-code/2024/aocutil"
)

//go:embed input*.txt
var Input embed.FS

type Graph map[string]map[string]struct{}

func (g Graph) AddEdge(from, to string) {
	if _, ok := g[from]; !ok {
		g[from] = make(map[string]struct{})
	}

	if _, ok := g[to]; !ok {
		g[to] = make(map[string]struct{})
	}
	g[from][to] = struct{}{}
	g[to][from] = struct{}{}
}

func (g Graph) Connected(from, to string) bool {
	_, ok := g[from][to]
	return ok
}

func ReadInput(inputName string) (graph Graph) {
	g := make(Graph)
	lines := aocutil.FSReadFile(Input, inputName)
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		nodes := strings.Split(line, "-")
		g.AddEdge(nodes[0], nodes[1])
	}
	return g
}

func Part1(g Graph) int {
	fccs := make(map[[3]string]struct{})

	for a := range g {
		if a[0] != 't' {
			continue
		}
		for b := range g[a] {
			for c := range g[b] {
				if g.Connected(a, c) {
					fcc := [3]string{a, b, c}
					sort.StringSlice(fcc[:]).Sort()
					fccs[fcc] = struct{}{}
				}
			}
		}
	}

	return len(fccs)
}

func GrowFCC(g Graph, component map[string]struct{}) []string {
onThrough:
	for {
		// Count number of edges to each neighbor of the FCC
		freqs := make(map[string]int)
		for node := range component {
			for neighbor := range g[node] {
				freqs[neighbor]++
			}
		}

		// If all nodes in FCC is connected to same neighbor,
		// add neighbor to FCC.
		for addNode := range freqs {
			if freqs[addNode] < len(component) {
				continue
			}
			component[addNode] = struct{}{}
			continue onThrough
		}

		// No more nodes to add, just return
		break onThrough // to the other side
	}

	// Return the FCC
	l := make([]string, 0, len(component))
	for node := range component {
		l = append(l, node)
	}
	sort.Strings(l)
	return l
}

func Part2(g Graph) string {
	max := 0
	maxFCC := []string{}
	for node := range g {
		fcc := GrowFCC(g, map[string]struct{}{node: {}})
		if len(fcc) > max {
			max = len(fcc)
			maxFCC = fcc
		}
	}
	return strings.Join(maxFCC, ",")
}

func Part12(inputName string) (int, string) {
	secrets := ReadInput(inputName)
	return Part1(secrets), Part2(secrets)
}
