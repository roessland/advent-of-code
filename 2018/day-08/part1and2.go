package main

import "fmt"
import "strconv"
import "io"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

type Node struct {
	Children []*Node
	Entries  []int
}

func ReadNum(r *bufio.Reader) int {
	word, _ := r.ReadString(' ')
	num, err := strconv.Atoi(strings.Trim(word, " \n"))
	Check(err)
	return num
}

func ReadNode(r *bufio.Reader) *Node {
	node := &Node{}
	numChildren := ReadNum(r)
	numEntries := ReadNum(r)
	for i := 0; i < numChildren; i++ {
		node.Children = append(node.Children, ReadNode(r))
	}
	for i := 0; i < numEntries; i++ {
		node.Entries = append(node.Entries, ReadNum(r))
	}
	return node
}

func SumEntries(node *Node) int {
	sum := Sum(node.Entries)
	for _, child := range node.Children {
		sum += SumEntries(child)
	}
	return sum
}

func Value(node *Node) int {
	if len(node.Children) == 0 {
		return Sum(node.Entries)
	}
	val := 0
	for _, childIdx := range node.Entries {
		if childIdx == 0 || childIdx-1 >= len(node.Children) {
			continue
		}
		val += Value(node.Children[childIdx-1])
	}
	return val
}

func main() {
	f, err := os.Open("input.txt")
	Check(err)
	r := bufio.NewReader(f)
	root := ReadNode(r)
	fmt.Println(SumEntries(root))
	fmt.Println(Value(root))
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
var _ = io.EOF

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
