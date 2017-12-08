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
		graph[name] = &Node{name, weight, children, -1}
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

func main() {
	graph := ReadInput()
	nodes := TopologicalSort(graph)
	fmt.Println(nodes[0].Name)
}
