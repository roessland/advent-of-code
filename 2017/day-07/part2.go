package main

import "fmt"
import "log"
import "sort"
import "strconv"
import "bufio"
import "strings"
import "os"

type Node struct {
	Name        string
	Weight      int
	TotalWeight int
	Children    []string
	VisitedTime int
}

func ReadInput() map[string]*Node {
	graph := map[string]*Node{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		name := parts[0]
		weight, err := strconv.Atoi(parts[1][1 : len(parts[1])-1])
		if err != nil {
			log.Fatal(err)
		}
		children := []string{}
		if len(parts) >= 3 {
			for _, child := range parts[3:len(parts)] {
				children = append(children, strings.Replace(child, ",", "", 1))
			}
		}
		graph[name] = &Node{name, weight, -1, children, -1}
	}
	return graph
}

func Dfs(graph map[string]*Node) {
	var currentTime int = 1
	var Visit func(map[string]*Node, string)
	Visit = func(graph map[string]*Node, name string) {
		if graph[name].VisitedTime > 0 {
			return
		}
		for _, child := range graph[name].Children {
			Visit(graph, child)
		}
		graph[name].VisitedTime = currentTime
		currentTime++
	}
	for name, _ := range graph {
		Visit(graph, name)
	}
}

func ToSlice(graph map[string]*Node) []*Node {
	slice := make([]*Node, 0, len(graph))
	for _, node := range graph {
		slice = append(slice, node)
	}
	return slice
}

func TopologicalSort(graph map[string]*Node) []*Node {
	Dfs(graph)
	nodes := ToSlice(graph)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].VisitedTime > nodes[j].VisitedTime
	})
	return nodes
}

func FindTotalWeights(graph map[string]*Node, node string) int {
	total := graph[node].Weight
	for _, child := range graph[node].Children {
		total += FindTotalWeights(graph, child)
	}
	graph[node].TotalWeight = total
	return total
}

func MedianTotalWeight(graph map[string]*Node, children []string) int {
	weights := make([]int, len(children))
	for i, child := range children {
		weights[i] = graph[child].TotalWeight
	}
	sort.Ints(weights)
	return weights[len(weights)/2]
}

func FindWrongNode(graph map[string]*Node, root string) {
	expected := MedianTotalWeight(graph, graph[root].Children)
	for _, child := range graph[root].Children {
		if graph[child].TotalWeight != expected {
			fmt.Println(child, "is", graph[child].Weight)
			fmt.Println("but should be",
				graph[child].Weight+expected-graph[child].TotalWeight)
			FindWrongNode(graph, child)
		}
	}
}

func main() {
	graph := ReadInput()
	nodes := TopologicalSort(graph)
	root := nodes[0].Name
	FindTotalWeights(graph, root)
	FindWrongNode(graph, root)
	fmt.Println("^^^^ last number is the answer ^^^^")
}
