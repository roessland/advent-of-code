package main

import "fmt"
import "strconv"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

type Node struct {
	name     string
	done     bool
	needs    map[string]bool
	children []string
}

func GetNext(g map[string]*Node) *Node {
	var ret *Node = nil
	lowestName := "["
	for _, node := range g {
		if node.done {
			continue
		}
		if len(node.needs) > 0 {
			continue
		}
		if node.name < lowestName {
			ret = node
			lowestName = node.name
		}
	}
	return ret
}

func main() {

	g := map[string]*Node{}

	buf, _ := ioutil.ReadFile("example.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		var step string
		var dependency string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dependency, &step)

		if g[step] == nil {
			g[step] = &Node{step, false, map[string]bool{dependency: true}, []string{}}
		} else {
			g[step].needs[dependency] = true
		}

		if g[dependency] == nil {
			g[dependency] = &Node{dependency, false, map[string]bool{}, []string{step}}
		} else {
			g[dependency].children = append(g[dependency].children, step)
		}
	}

	for len(g) > 0 {
		next := GetNext(g)
		if next == nil {
			break
		}
		for _, childName := range next.children {
			delete(g[childName].needs, next.name)
		}
		next.done = true
		fmt.Printf("%s", next.name)
	}

}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
