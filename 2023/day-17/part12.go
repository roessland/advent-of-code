package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"math"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"golang.org/x/exp/constraints"
)

func main() {
	m := ReadInput()
	fmt.Println("Part 1:", solve(&m, 0, 3))
	fmt.Println("Part 2:", solve(&m, 4, 10))
}

type HeapItem[T any, V cmp.Ordered] struct {
	Item  T
	Value V
}

type PriorityQueue[T any, V cmp.Ordered] struct {
	heap ArrayHeap[T, V]
}

// ArrayHeap implements heap.Interface.
type ArrayHeap[T any, V cmp.Ordered] ([]HeapItem[T, V])

var _ heap.Interface = new(ArrayHeap[int, float64])

func (ah ArrayHeap[T, V]) Len() int {
	return len(ah)
}

func (ah ArrayHeap[T, V]) Less(i int, j int) bool {
	return ah[i].Value < ah[j].Value
}

func (ah *ArrayHeap[T, V]) Pop() any {
	hItem := (*ah)[len(*ah)-1]
	*ah = (*ah)[:len(*ah)-1]
	return hItem
}

func (ah *ArrayHeap[T, V]) Push(x any) {
	*ah = append(*ah, x.(HeapItem[T, V]))
}

func (ah ArrayHeap[T, V]) Swap(i int, j int) {
	ah[i], ah[j] = ah[j], ah[i]
}

func NewPriorityQueue[T any, V cmp.Ordered]() *PriorityQueue[T, V] {
	arr := ArrayHeap[T, V](make([]HeapItem[T, V], 0))
	return &PriorityQueue[T, V]{
		heap: arr,
	}
}

func (pq *PriorityQueue[T, V]) Push(item T, value V) {
	hItem := HeapItem[T, V]{Item: item, Value: value}
	heap.Push(&pq.heap, hItem)
}

func (pq *PriorityQueue[T, V]) Pop() T {
	return heap.Pop(&pq.heap).(HeapItem[T, V]).Item
}

func (pq *PriorityQueue[T, V]) Len() int {
	return len(pq.heap)
}

type NodeID int

type NodePair struct {
	A, B NodeID
}

type Graph[Dist constraints.Float | constraints.Integer] interface {
	Neighbors(id NodeID) []NodeID
	Distance(a, b NodeID) Dist
	Infinity() Dist
}

// Finds the shortest distance from the start node to the end node.
func Dijkstra[Dist constraints.Float | constraints.Integer](g Graph[Dist], start NodeID, isEnd func(NodeID) bool) (dist map[NodePair]Dist, prev map[NodeID]NodeID) {
	dist = make(map[NodePair]Dist)
	prev = make(map[NodeID]NodeID)
	pq := NewPriorityQueue[NodeID, Dist]()
	pq.Push(start, Dist(0))
	for pq.Len() > 0 {
		currentID := pq.Pop()
		currentDist := dist[NodePair{start, currentID}]

		if isEnd(currentID) {
			return dist, prev
		}

		for _, neighborID := range g.Neighbors(currentID) {

			currNeighborDist, ok := dist[NodePair{start, neighborID}]
			if !ok {
				currNeighborDist = g.Infinity()
			}

			edgeCost := g.Distance(currentID, neighborID)
			altNeighborDist := currentDist + edgeCost

			if altNeighborDist < currNeighborDist {
				dist[NodePair{start, neighborID}] = altNeighborDist
				prev[neighborID] = currentID
				pq.Push(neighborID, altNeighborDist)
			}
		}
	}

	panic("unreachable")
}

func ReadInput() Map {
	input := aocutil.ReadLinesAsBytes("input.txt")
	for y := range input {
		for x := range input[y] {
			input[y][x] -= '0'
		}
	}
	return Map{Height: len(input), Width: len(input[0]), EntryLoss: input}
}

type Dirs int

const (
	Right    Dirs = 1
	Up       Dirs = 2
	Left     Dirs = 4
	Down     Dirs = 8
	Nowhere  Dirs = 16
	Anywhere Dirs = Right | Up | Left | Down
)

func (d Dirs) Opposite() Dirs {
	switch d {
	case Right:
		return Left
	case Up:
		return Down
	case Left:
		return Right
	case Down:
		return Up
	default:
		panic("no opposite")
	}
}

func (d Dirs) AsList() []Dirs {
	ds := make([]Dirs, 0)
	if d == Nowhere {
		return nil
	}
	if d&Right > 0 {
		ds = append(ds, Right)
	}
	if d&Up > 0 {
		ds = append(ds, Up)
	}
	if d&Left > 0 {
		ds = append(ds, Left)
	}
	if d&Down > 0 {
		ds = append(ds, Down)
	}
	return ds
}

func dx(d Dirs) int {
	switch d {
	case Right:
		return 1
	case Left:
		return -1
	default:
		return 0
	}
}

func dy(d Dirs) int {
	switch d {
	case Up:
		return -1
	case Down:
		return 1
	default:
		return 0
	}
}

type Map struct {
	EntryLoss [][]byte
	Height    int
	Width     int
}

type ProblemGraph struct {
	m              *Map
	nodes          map[NodeID]ProblemNode
	ids            map[ProblemNode]NodeID
	nextID         NodeID
	minConsecutive int
	maxConsecutive int
}

type ProblemNode struct {
	Y, X        int
	Dir         Dirs
	CanGo       Dirs
	Consecutive int
}

func NewProblemGraph(m *Map, minConsecutive, maxConsecutive int) *ProblemGraph {
	return &ProblemGraph{
		nodes:          make(map[NodeID]ProblemNode),
		ids:            make(map[ProblemNode]NodeID),
		m:              m,
		minConsecutive: minConsecutive,
		maxConsecutive: maxConsecutive,
	}
}

func (g *ProblemGraph) Infinity() int {
	return 1000000000
}

func (g *ProblemGraph) Distance(a, b NodeID) int {
	by, bx := g.nodes[b].Y, g.nodes[b].X
	if by < 0 || by >= g.m.Height || bx < 0 || bx >= g.m.Width {
		return g.Infinity()
	}
	return int(g.m.EntryLoss[by][bx])
}

func (g *ProblemGraph) Neighbors(id NodeID) []NodeID {
	ns := make([]NodeID, 0)

	c := g.nodes[id]

	// Continue in the same direction
	if c.Dir != Nowhere && c.Consecutive < g.maxConsecutive {
		ny := c.Y + dy(c.Dir)
		nx := c.X + dx(c.Dir)
		n := ProblemNode{Y: ny, X: nx, Dir: c.Dir, CanGo: c.CanGo, Consecutive: c.Consecutive + 1}
		ns = append(ns, g.getID(n))
	}

	// Turn
	if c.Consecutive >= g.minConsecutive {
		for _, dir := range c.CanGo.AsList() {
			ny := c.Y + dy(dir)
			nx := c.X + dx(dir)
			canGo := Anywhere - dir - dir.Opposite()
			n := ProblemNode{Y: ny, X: nx, Dir: dir, CanGo: canGo, Consecutive: 1}
			ns = append(ns, g.getID(n))
		}
	}

	return ns
}

func (g *ProblemGraph) getID(n ProblemNode) NodeID {
	if id, ok := g.ids[n]; ok {
		return id
	}
	id := g.nextID
	g.nextID++

	g.ids[n] = id
	g.nodes[id] = n
	return id
}

func solve(m *Map, minConsecutive, maxConsecutive int) int {
	g := NewProblemGraph(m, minConsecutive, maxConsecutive)
	start := ProblemNode{Y: 0, X: 0, CanGo: Down | Right, Consecutive: maxConsecutive}
	startID := g.getID(start)
	isEnd := func(id NodeID) bool {
		n, ok := g.nodes[id]
		if !ok {
			panic("can't find node")
		}
		return n.Y == m.Height-1 && n.X == m.Width-1
	}
	dists, _ := Dijkstra[int](g, startID, isEnd)

	minDist := math.MaxInt
	for fromToPair, dist := range dists {
		if fromToPair.A != startID || !isEnd(fromToPair.B) {
			continue
		}
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}
