package aocutil

import (
	"fmt"
	"os"
	"os/exec"
)

func VisualizeIntGraph(g map[int]map[int]bool) {
	f, err := os.Create("graph.dot")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(f, "digraph G {\n")
	for i, edges := range g {
		for j := range edges {
			fmt.Fprintf(f, "%d -> %d\n", i, j)
		}
	}
	fmt.Fprintf(f, "}\n")
	f.Close()
	output, err := exec.Command("dot", "-Tpng", "graph.dot", "-o", "graph.png").CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		panic(err)
	}

	output2, err2 := exec.Command("open", "graph.png").CombinedOutput()
	fmt.Println(string(output2))
	if err2 != nil {
		panic(err2)
	}
}
