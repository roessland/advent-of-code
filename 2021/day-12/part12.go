package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)



func main() {
	t0 := time.Now()
	g := ReadInput()
	part1(g)
	part2(g)
	fmt.Println(time.Since(t0))

}

func part1(g Graph) {
	visited := map[string]int{}
	paths := search(g, "start", visited, false)
	fmt.Println("Part 1:", paths)
}

func part2(g Graph) {
	visited := map[string]int{}
	paths := search(g, "start", visited, true)
	fmt.Println("Part 2:", paths)
}

func search(g Graph, fromId string, visited map[string]int, token bool) int {
	if fromId == "end" {
		return 1
	}
	if IsSmall(fromId) {
		visited[fromId]++
	}
	var paths int
	for _, neighborId := range g[fromId] {
		if neighborId == "start" {
			continue
		}
		if visited[neighborId] == 0 {
			// Visit without using token
			paths += search(g, neighborId, visited, token)
		} else if IsSmall(neighborId) && token {
			// Visit using token
			paths += search(g, neighborId, visited, false)

		}
	}
	if IsSmall(fromId) {
		visited[fromId]--
	}
	return paths
}

type Graph map[string][]string

func (g Graph) AddNode(id string) {
	if g[id] == nil {
		g[id] = make([]string, 0)
	}
}

func (g Graph) AddEdge(nodeId1 string, nodeId2 string) {
	g[nodeId1] = append(g[nodeId1], nodeId2)
	g[nodeId2] = append(g[nodeId2], nodeId1)
}

func IsSmall(nodeId string) bool {
	return strings.ToLower(nodeId) == nodeId
}

func ReadInput() Graph {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	g := make(Graph)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		g.AddNode(nodes[0])
		g.AddNode(nodes[1])
		g.AddEdge(nodes[0], nodes[1])
	}
	return g
}