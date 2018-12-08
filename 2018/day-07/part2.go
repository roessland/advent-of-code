package main

import "fmt"
import "sort"
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
	timeLeft int
	active   bool
}

func GetInitialTime(nodeName string) int {
	t := nodeName[0] - 'A' + 1 + 60 // FIX
	return int(t)
}

func GetNexts(g map[string]*Node) []*Node {
	ret := []*Node{}
	for _, node := range g {
		if node.done {
			continue
		}
		if len(node.needs) > 0 {
			continue
		}
		if node.timeLeft == 0 {
			continue
		}
		ret = append(ret, node)
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].active && ret[j].active {
			return ret[i].name < ret[j].name
		}
		return ret[i].active

	})
	if len(ret) < 5 { // FIX
		return ret[0:len(ret)]
	} else {
		return ret[0:5] // FIX
	}
}

func main() {

	g := map[string]*Node{}

	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		var step string
		var dependency string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dependency, &step)

		if g[step] == nil {
			T := GetInitialTime(step)
			g[step] = &Node{step, false, map[string]bool{dependency: true}, []string{}, T, false}
		} else {
			g[step].needs[dependency] = true
		}

		if g[dependency] == nil {
			T := GetInitialTime(dependency)
			g[dependency] = &Node{dependency, false, map[string]bool{}, []string{step}, T, false}
		} else {
			g[dependency].children = append(g[dependency].children, step)
		}
	}

	sec := 0
	for len(g) > 0 {
		nexts := GetNexts(g)
		if len(nexts) == 0 {
			break
		}
		if len(nexts) >= 1 {
			Work(g, nexts[0])
		}
		if len(nexts) >= 2 {
			Work(g, nexts[1])
		}
		if len(nexts) >= 3 {
			Work(g, nexts[2])
		}
		if len(nexts) >= 4 {
			Work(g, nexts[3])
		}
		if len(nexts) >= 5 {
			Work(g, nexts[4])
		}
		sec++
	}
	fmt.Println("\nSeconds: ", sec)

}

func Work(g map[string]*Node, node *Node) {
	if node == nil {
		return
	}
	node.active = true
	node.timeLeft--
	if node.timeLeft == 0 {
		fmt.Printf("%s", node.name)
		for _, childName := range node.children {
			delete(g[childName].needs, node.name)
		}
		node.done = true
		node.active = false
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
