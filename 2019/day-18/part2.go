package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

const TileWall rune = '#'
const TileEmpty rune = '.'
const TileStart rune = '@'
const Unreachable int = math.MaxInt32

func ReadInput(filename string) [][]rune {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	out := make([][]rune, 0)
	line := make([]rune, 0)
	for {
		r, _, err := reader.ReadRune()
		if r == '\r' {
			continue
		}
		if err == io.EOF {
			if len(line) > 0 {
				out = append(out, line)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		} else if r == '\n' {
			out = append(out, line)
			line = nil
		} else {
			line = append(line, r)
		}
	}
	return out
}

func PrintMap(tiles [][]rune) {
	for _, line := range tiles {
		fmt.Printf("%s\n", string(line))
	}
}

type Pos struct {
	I, J int
}

type Edge struct {
	To     *Node
	Weight int
}

type Node struct {
	Edges []*Edge
	Pos   Pos
	Tile  rune
}

func (n *Node) DotNodeName() string {
	return fmt.Sprintf(`%c\n(%d,%d)`, n.Tile, n.Pos.I, n.Pos.J)
}

func CreateWorld(input [][]rune) map[Pos]*Node {
	nodeAtPos := map[Pos]*Node{}
	height := len(input)
	width := len(input[0])
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			pos := Pos{I: i, J: j}
			// make da node
			n := Node{
				Edges: make([]*Edge, 0),
				Pos:   pos,
				Tile:  input[i][j],
			}
			if n.Tile != TileWall {
				nodeAtPos[pos] = &n
			}
		}
	}
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			this := nodeAtPos[Pos{I: i, J: j}]
			if this == nil {
				continue
			}
			east := nodeAtPos[Pos{I: i, J: j + 1}]
			west := nodeAtPos[Pos{I: i, J: j - 1}]
			south := nodeAtPos[Pos{I: i + 1, J: j}]
			north := nodeAtPos[Pos{I: i - 1, J: j}]
			for _, dir := range []*Node{east, west, south, north} {
				if dir == nil {
					continue
				}
				if dir.Tile != TileWall {
					this.Edges = append(this.Edges, &Edge{dir, 1})
				}
			}
		}
	}
	return nodeAtPos
}

func WorldToDot(world map[Pos]*Node, filename string) {
	var lines []string
	lines = append(lines, "digraph G {")
	for _, node := range world {
		lines = append(lines, fmt.Sprintf(`"%s" [`, node.DotNodeName()))
		lines = append(lines, fmt.Sprintf(`    pos = "%d,%d!"`, node.Pos.J, -node.Pos.I))
		lines = append(lines, "]")
		for _, edge := range node.Edges {
			line := fmt.Sprintf(`"%s" -> "%s"`, node.DotNodeName(), edge.To.DotNodeName())
			lines = append(lines, line)
		}
	}

	lines = append(lines, "}")
	err := ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func IsDoor(tile rune) bool {
	return 'A' <= tile && tile <= 'Z'
}

func IsKey(tile rune) bool {
	return 'a' <= tile && tile <= 'z'
}

func HasKeyToDoor(tile rune, keySet map[rune]struct{}) bool {
	if tile < 'A' || 'Z' < tile {
		return false
	} else {
		_, hasKey := keySet[tile-'A'+'a']
		return hasKey
	}
}

func GetStarts(world map[Pos]*Node) []*Node {
	var starts []*Node
	for _, n := range world {
		if n.Tile == TileStart {
			starts = append(starts, n)
		}
	}
	return starts
}

type State struct {
	Nodes     []*Node
	KeySet    map[rune]struct{}
	Distance  int
	PrevState *State
	index     int // for heap
}

type States map[string]*State

func (ss States) GetOrCreate(nodes []*Node, keySet map[rune]struct{}) *State {
	id := Id(nodes, keySet)
	if ss[id] == nil {
		ss[id] = NewStateNode(nodes, keySet)
	}
	return ss[id]
}

func (ss States) Get(id string) *State {
	n, ok := ss[id]
	if !ok {
		panic("tried to get non-existing statenode")
	}
	return n
}

func Id(nodes []*Node, keySet map[rune]struct{}) string {
	var keys []rune
	for k, _ := range keySet {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	pos := ""
	for _, n := range nodes {
		if n == nil {
			fmt.Println(nodes)
			panic("cant id nodes where one is nil")
		}
		pos += fmt.Sprintf("-(%d,%d)", n.Pos.I, n.Pos.J)
	}
	return fmt.Sprintf("%s-%s", string(keys), pos)
}

func (s *State) Id() string {
	return Id(s.Nodes, s.KeySet)
}

func NewStateNode(nodes []*Node, keySet map[rune]struct{}) *State {
	ret := &State{
		Nodes:    nodes,
		KeySet:   keySet,
		Distance: math.MaxInt32,
	}
	return ret
}

func StatePriorityQueuePop(pq map[*State]struct{}) *State {
	var minState *State
	minDist := Unreachable
	for state, _ := range pq {
		if state.Distance < minDist {
			minState = state
			minDist = state.Distance
		}
	}
	delete(pq, minState)
	return minState
}

func NodePriorityQueuePop(distanceTo map[*Node]int, pq map[*Node]struct{}) *Node {
	var minNode *Node
	minDist := Unreachable
	for node, _ := range pq {
		if distanceTo[node] < minDist {
			minNode = node
			minDist = distanceTo[node]
		}
	}
	delete(pq, minNode)
	return minNode
}

// WithKey creates a copy of a keyset with an additional key added, but only if it's a key.
func WithKey(keySet map[rune]struct{}, maybeKey rune) map[rune]struct{} {
	if IsKey(maybeKey) {
		newKeySet := map[rune]struct{}{}
		newKeySet[maybeKey] = struct{}{}
		for key, _ := range keySet {
			newKeySet[key] = struct{}{}
		}
		return newKeySet
	}
	return keySet
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *State, distance int) {
	item.Distance = distance
	heap.Fix(pq, item.index)
}

// GetOrCreate gets node at some position, creating it if necessary.
func GetOrCreate(world map[Pos]*Node, pos Pos, tile rune) *Node {
	if world[pos] == nil {
		world[pos] = &Node{
			Edges: nil,
			Pos:   pos,
			Tile:  tile,
		}
	}
	return world[pos]
}

func Part1(world map[Pos]*Node) {
	targetKeyCount := CountKeys(world)
	states := make(States)
	initial := states.GetOrCreate(GetStarts(world), make(map[rune]struct{}))
	initial.Distance = 0

	pq := make(PriorityQueue, 1)
	pq[0] = initial

	maxKeySetSize := 0
	var finalState *State
	for len(pq) > 0 {
		currState := heap.Pop(&pq).(*State)
		if len(currState.KeySet) == targetKeyCount {
			finalState = currState
			break
		}

		// Just some output to print progress
		if len(currState.KeySet) > maxKeySetSize {
			maxKeySetSize = len(currState.KeySet)
			fmt.Printf("%d/%d ", maxKeySetSize, targetKeyCount)
		}

		// Future state, with only one robot moved
		currNodes := currState.Nodes
		for r, _ := range currNodes {
			// Find edges
			currNode := currNodes[r]
			for _, edge := range currNode.Edges {
				if IsDoor(edge.To.Tile) && !HasKeyToDoor(edge.To.Tile, currState.KeySet) {
					continue
				}
				toKeySet := WithKey(currState.KeySet, edge.To.Tile)
				nodesCopy := make([]*Node, len(currNodes))
				copy(nodesCopy, currNodes)
				nodesCopy[r] = edge.To
				toState := states.GetOrCreate(nodesCopy, toKeySet)

				// If going via this state is faster, update and add to queue
				if currState.Distance+edge.Weight < toState.Distance {
					toState.Distance = currState.Distance + edge.Weight
					toState.PrevState = currState
					heap.Push(&pq, toState)
				}
			}
		}
	}

	if finalState == nil {
		log.Fatal("couldn't find final state")
	}
	fmt.Println("\nPart 1:", finalState.Distance)
}

// CountKeys finds our goal.
func CountKeys(w map[Pos]*Node) int {
	num := 0
	for _, n := range w {
		if IsKey(n.Tile) {
			num++
		}
	}
	return num
}

// Preprocess creates a new world with all the dots omitted.
// This gives a ~10-50x speedup.
func Preprocess(oldWorld map[Pos]*Node) map[Pos]*Node {
	newWorld := map[Pos]*Node{}
	// Foreach non-dot tile
	for _, oldNode := range oldWorld {
		if oldNode.Tile == TileEmpty {
			continue
		}
		distanceTo := map[*Node]int{oldNode: 0}
		// Not implemented using heap. O(n) pop.
		pq := map[*Node]struct{}{oldNode: {}}
		// Find everything reachable without passing through objects
		for len(pq) > 0 {
			node := NodePriorityQueuePop(distanceTo, pq)
			for _, e := range node.Edges {
				_, hasDistanceTo := distanceTo[e.To]
				if !hasDistanceTo {
					distanceTo[e.To] = Unreachable
				}
				if distanceTo[node] + e.Weight < distanceTo[e.To] {
					distanceTo[e.To] = distanceTo[node] + e.Weight
					if e.To.Tile == TileEmpty {
						pq[e.To] = struct{}{}
					}
				}
			}
		}

		// Foreach reachable non-empty edge, add edge
		newFrom := GetOrCreate(newWorld, oldNode.Pos, oldNode.Tile)
		for oldTo, dist := range distanceTo {
			if oldTo == oldNode {
				continue
			}
			if oldTo.Tile == TileEmpty {
				continue
			}
			newTo := GetOrCreate(newWorld, oldTo.Pos, oldTo.Tile)
			newFrom.Edges = append(newFrom.Edges, &Edge{
				To:     newTo,
				Weight: dist,
			})
		}
	}
	return newWorld
}

func main() {
	input := ReadInput("input-part2.txt")
	w := CreateWorld(input)
	w = Preprocess(w)
	Part1(w) // 4248 too high
}
