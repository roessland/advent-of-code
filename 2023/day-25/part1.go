package main

import (
	"fmt"
	"strings"
	"time"

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

	t0 := time.Now()
	for i := 0; i < 100000; i++ {
		res := karger.Karger(g)

		if len(res.Edges) == 3 && res.SizeA > 1 && res.SizeB > 1 {
			ans := res.SizeA * res.SizeB
			fmt.Println("---")
			fmt.Println("Answer: ", ans)
			fmt.Println("Time: ", time.Since(t0))
			fmt.Println("(Took ", i, "iterations)")
			break
		}
	}
}
