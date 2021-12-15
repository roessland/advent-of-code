package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/roessland/gopkg/digraph"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	g, I1, J1 := ReadInput()
	I5, J5 := I1*5, J1*5

	dist, _ := digraph.Dijkstra(g, 0)
	fmt.Println("Part 2:", dist[(J5-1)*(I5)+(I5)-1]) // 3012

	// Remove edges going out from top left area so the part 1 doesn't go out of bounds.
	// j5*I5 + i5
	for id := 0; id < len(g.Nodes); id++ {
		if id / I5 >= J1 || id % J5 >= I1 {
			// We're outside top left area already.
			continue
		}
		for edgeNo, edge := range g.Nodes[id].Neighbors {
			if edge.To / I5 >= J1 || edge.To % J5 >= I1 {
				g.Nodes[id].Neighbors[edgeNo].Weight = math.MaxFloat64
			}
		}
	}
	dist, _ = digraph.Dijkstra(g, 0)
	fmt.Println("Part 1:", dist[(J1-1)*I5+I1-1]) // 811

	fmt.Println(time.Since(t0)) // 160 ms
}


func ReadInput() (graph digraph.Graph, I1, J1 int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var cave [][]int
	for scanner.Scan() {
		var nums []int
		for _, c := range scanner.Text() {
			nums = append(nums, int(c-'0'))
		}
		cave = append(cave, nums)
	}

	J1, I1 = len(cave), len(cave[0])
	J5, I5 := J1*5, I1*5
	graph = digraph.Graph{Nodes: make([]digraph.Node, I5*J5)}

	getID := func(j5, i5 int) (int, error) {
		if j5 < 0 || j5 >= J5 || i5 < 0 || i5 >= I5 {
			return -1, errors.New("oob")
		}
		return j5*I5 + i5, nil
	}

	getWeight := func(j5, i5 int) float64 {
		j1, i1 := j5%J1, i5%I1
		return float64((cave[j1][i1] + j5/J1 + i5/I1 - 1) % 9 + 1)
	}

	for j := 0; j < J5; j++ {
		for i := 0; i < I5; i++ {
			id, _ := getID(j, i)
			if neighbor, err := getID(j-1, i); err == nil {
				graph.Nodes[id].Neighbors = append(graph.Nodes[id].Neighbors, digraph.Edge{
					To:     neighbor,
					Weight: getWeight(j-1, i),
				})
			}
			if neighbor, err := getID(j+1, i); err == nil {
				graph.Nodes[id].Neighbors = append(graph.Nodes[id].Neighbors, digraph.Edge{
					To:     neighbor,
					Weight: getWeight(j+1, i),
				})
			}
			if neighbor, err := getID(j, i-1); err == nil {
				graph.Nodes[id].Neighbors = append(graph.Nodes[id].Neighbors, digraph.Edge{
					To:     neighbor,
					Weight: getWeight(j, i-1),
				})
			}
			if neighbor, err := getID(j, i+1); err == nil {
				graph.Nodes[id].Neighbors = append(graph.Nodes[id].Neighbors, digraph.Edge{
					To:     neighbor,
					Weight: getWeight(j, i+1),
				})
			}

		}
	}
	return
}
