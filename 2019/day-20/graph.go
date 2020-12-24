package main

import "math"

// GRAPH INTERFACES

type Iterator interface {
	Next() bool
}

type Node interface {
	ID() string
}

type Nodes interface {
	Iterator
	Node() Node
}

type Edge interface {
	From() string
	To() string
	Weight() int
}

type Edges interface {
	Iterator
	Edge() Edge
}

type Graph interface {
	AddNode(Node)
	Edge(string, string) Edge
	From(string) []Node
	HasEdgeBetween(string, string) bool
	NewNode(string) Node
	NewEdge(string, string, int) Edge
	Node(id string) Node
	Nodes() []Node
	Weight(string, string) int
	AddEdge(Edge)
	GetEdge(string, string) *Edge
}

func NewGraph() Graph {
	g := &simpleGraph{}
	g.nodes = make(map[string]Node)
	g.edges = make(map[string]map[string]Edge)
	g.absent = math.MaxInt32
	return g
}

// NODE IMPLEMENTATION

type simpleNode struct {
	id string
}

func (n *simpleNode) ID() string {
	return n.id
}

var _ Node = &simpleNode{}

// EDGE IMPLEMENTATION

type simpleEdge struct {
	from, to string
	weight   int
}

func (e *simpleEdge) From() string {
	return e.from
}

func (e *simpleEdge) To() string {
	return e.to
}

func (e *simpleEdge) Weight() int {
	return e.weight
}

var _ Edge = &simpleEdge{}

// GRAPH IMPLEMENTATION

type simpleGraph struct {
	nodes   map[string]Node
	edges   map[string]map[string]Edge
	absent  int
}

func (g *simpleGraph) AddNode(n Node) {
	id := n.ID()
	if g.nodes[id] != nil {
		panic("node already exists")
	}
	g.edges[id] = make(map[string]Edge)
	g.nodes[id] = n
}

func (g *simpleGraph) Edge(uid, vid string) Edge {
	if g.edges[uid] == nil {
		return nil
	}
	e, ok := g.edges[uid][vid]
	if ok {
		return e
	} else {
		return &simpleEdge{
			from:   uid,
			to:     uid,
			weight: g.absent,
		}
	}
}

func (g *simpleGraph) Node(id string) Node {
	n, ok := g.nodes[id]
	if !ok {
		panic("no such node)")
	}
	return n
}

func (g *simpleGraph) Nodes() []Node {
	var ns []Node
	for _, n := range g.nodes {
		ns = append(ns, n)
	}
	return ns
}

func (g *simpleGraph) From(id string) []Node {
	var ns []Node
	for _, e := range g.edges[id] {
		ns = append(ns, g.nodes[e.To()])
	}
	return ns
}

func (g *simpleGraph) HasEdgeBetween(uid, vid string) bool {
	_, ok := g.edges[uid][vid]
	return ok
}

func (g *simpleGraph) NewNode(id string) Node {
	return &simpleNode{id}
}

func (g *simpleGraph) NewEdge(uid, vid string, weight int) Edge {
	return &simpleEdge{from: uid, to: vid, weight: weight}
}

func (g *simpleGraph) Weight(uid, vid string) int {
	_, uok := g.nodes[uid]
	_, vok := g.nodes[vid]
	if !uok {
		panic("Weight: uid doesn't exist")
	}
	if !vok {
		panic("Weight: vid doesn't exist")
	}
	e, ok := g.edges[uid][vid]
	if ok {
		return e.Weight()
	}
	return g.absent
}

func (g *simpleGraph) AddEdge(e Edge) {
	if g.nodes[e.From()] == nil {
		panic("AddEdge: uid doesn't exist")
	}
	if g.nodes[e.To()] == nil {
		panic("AddEdge: vid doesn't exist")
	}
	if g.edges[e.From()][e.To()] != nil {
		panic("AddEdge: edge already exists")
	}
	g.edges[e.From()][e.To()] = e
}

func (g *simpleGraph) GetEdge(from, to string) *Edge {
	if g.nodes[from] == nil {
		panic("GetEdge: uid doesn't exist")
	}
	if g.nodes[to] == nil {
		panic("GetEdge: vid doesn't exist")
	}
	e, ok := g.edges[from][to]
	if ok {
		return &e
	}
	return nil
}

/*
GetEdge(string, string) *Edge
*/
