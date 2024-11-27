package main

import (
	"fmt"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/roessland/gopkg/graph/karger"
)

func main() {
	vertices := map[string]int{}
	nextID := 0
	getID := func(name string) int {
		if id, ok := vertices[name]; ok {
			return id
		} else {
			id := nextID
			vertices[name] = id
			nextID++
			return id
		}
	}

	edges := []karger.Edge{}
	lines := aocutil.ReadLines("input.txt")
	for _, line := range lines {
		parts := strings.Split(strings.Replace(line, ":", "", 1), " ")
		for _, part := range parts[1:] {
			edges = append(edges, karger.Edge{A: getID(parts[0]), B: getID(part)})
		}
	}

	g := karger.NewUnweightedGraph(len(vertices))
	g.Edges = edges

	greatest := 0
	for i := 0; i < 1000000; i++ {
		result := karger.Karger(g)
		alt := result.SizeA * result.SizeB
		if len(result.Edges) != 3 {
			continue
		}
		if alt > greatest {
			greatest = alt
			fmt.Println(greatest, result)
		}
	}
}
