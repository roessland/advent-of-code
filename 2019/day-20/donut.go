package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)



// NODE IMPLEMENTATION

type donutNode struct {
	id string
	level int
	pos Pos
}

func (g *donutGraph) IsOuter(pos Pos) bool {
	return pos.I < 4 || pos.J < 4 || pos.I > g.I-4 || pos.J > g.J-4
}

func (n *donutNode) ID() string {
	return n.id
}

var _ Node = &donutNode{}

// EDGE IMPLEMENTATION

type donutEdge struct {
	from, to string
	weight   int
}

func (e *donutEdge) From() string {
	return e.from
}

func (e *donutEdge) To() string {
	return e.to
}

func (e *donutEdge) Weight() int {
	return e.weight
}

var _ Edge = &donutEdge{}

// GRAPH IMPLEMENTATION

type donutGraph struct {
	nodes   map[string]Node
	edges   map[string]map[string]Edge
	absent  int
	tiles []string
	I, J int
	portalNameByPos map[Pos]string
	outerPortals map[string]Pos
	innerPortals map[string]Pos
}

func NewDonutGraph(tiles []string) (Graph, Node, Node) {
	g := &donutGraph{}
	g.nodes = make(map[string]Node)
	g.edges = make(map[string]map[string]Edge)
	g.absent = math.MaxInt32
	g.tiles = tiles
	g.I = len(tiles)
	g.J = len(tiles[0])
	g.portalNameByPos = make(map[Pos]string)
	g.outerPortals = make(map[string]Pos)
	g.innerPortals = make(map[string]Pos)
	g.findPortals()
	aaPos := g.outerPortals["AA"]
	zzPos := g.outerPortals["ZZ"]
	aaNode := g.Node(fmt.Sprintf("0-%d-%d", aaPos.I, aaPos.J))
	zzNode := g.Node(fmt.Sprintf("0-%d-%d", zzPos.I, zzPos.J))
	return g, aaNode, zzNode
}

func (g *donutGraph) AddNode(n Node) {
	id := n.ID()
	if g.nodes[id] != nil {
		fmt.Println("id: ", id)
		panic("node already exists")
	}
	g.edges[id] = make(map[string]Edge)
	g.nodes[id] = n
}

func (g *donutGraph) Edge(uid, vid string) Edge {
	if g.edges[uid] == nil {
		return nil
	}
	e, ok := g.edges[uid][vid]
	if ok {
		return e
	} else {
		return &donutEdge{
			from:   uid,
			to:     uid,
			weight: g.absent,
		}
	}
}

// %d-%d-%d
// level-i-j
func newDonutNodeFromId(id string) Node {
	parts := strings.Split(id, "-")
	var n donutNode
	n.level, _ = strconv.Atoi(parts[0])
	n.pos.I, _ = strconv.Atoi(parts[1])
	n.pos.J, _ = strconv.Atoi(parts[2])
	n.id = id

	return &n
}

func findPortal(tiles []string, i1, j1 int) (name string, pos Pos) {
	c1 := tiles[i1][j1]
	cs := []byte{c1}
	cpos := []Pos{{i1, j1}}
	if !isLetter(c1) {
		panic("not a letter")
	}

	// See if dot is near this letter
	dotPos := findDot(tiles, i1, j1)

	var i2, j2 int
	if i1 > 0 && isLetter(tiles[i1-1][j1]) {
		i2, j2 = i1-1, j1
	} else if i1 < len(tiles) -1 && isLetter(tiles[i1+1][j1]) {
		i2, j2 = i1+1, j1
	} else if j1 > 0 && isLetter(tiles[i1][j1-1]) {
		i2, j2 = i1, j1-1
	} else if j1 < len(tiles[0]) - 1 && isLetter(tiles[i1][j1+1]) {
		i2, j2 = i1, j1+1
	}
	c2 := tiles[i2][j2]
	cs = append(cs, c2)
	cpos = append(cpos, Pos{i2, j2})

	// See if dot is near this letter
	if dotPos == nil {
		dotPos = findDot(tiles, i2, j2)
	}
	if dotPos == nil {
		panic("found no dot for this portal")
	}

	sort.Slice(cs, func(i,j int)bool {
		return (cpos[i].I + cpos[i].J) < (cpos[j].I + cpos[j].J)
	})
	name = string(cs)

	return name, *dotPos
}

func (g *donutGraph) findPortals() {
	for i := 0; i < g.I; i++ {
		for j := 0; j < g.J; j++ {
			if !isLetter(g.tiles[i][j]) {
				continue
			}
			portalName, dotPos := findPortal(g.tiles, i, j)
			g.portalNameByPos[dotPos] = portalName
			if g.IsOuter(dotPos) {
				g.outerPortals[portalName] = dotPos
			} else {
				g.innerPortals[portalName] = dotPos
			}
		}
	}
}

// Node gets a node, creating it if it doesn't already exist.
func (g *donutGraph) Node(id string) Node {
	n, ok := g.nodes[id]
	if !ok {
		n = newDonutNodeFromId(id)
		g.AddNode(n)
	}
	return n
}

// There are infinitely many nodes so this doesn't make sense.
func (g *donutGraph) Nodes() []Node {
	panic("not implemented")
}

// From gets neighbors of node with a given ID.
// Lazy graph.
func (g *donutGraph) From(id string) []Node {
	var ns []Node
	n := g.Node(id).(*donutNode)

	// Neighboring tiles in same level
	for _, pos := range []Pos{n.pos.Down(), n.pos.Up(), n.pos.Left(), n.pos.Right()} {
		if g.tiles[pos.I][pos.J] == '.' {
			ns = append(ns, g.Node(fmt.Sprintf("%d-%d-%d", n.level, pos.I, pos.J)))
		}
	}

	// Portals
	if portalName, ok := g.portalNameByPos[n.pos]; ok {
		if portalName == "AA" || portalName == "ZZ" {
			// cant go there
		} else if g.IsOuter(n.pos) && n.level > 0 {
			dstPos := g.innerPortals[portalName]
			ns = append(ns, g.Node(fmt.Sprintf("%d-%d-%d", n.level-1, dstPos.I, dstPos.J)))
		} else if !g.IsOuter(n.pos) {
			dstPos := g.outerPortals[portalName]
			ns = append(ns, g.Node(fmt.Sprintf("%d-%d-%d", n.level+1, dstPos.I, dstPos.J)))
		}
	}

	return ns
}

func (g *donutGraph) HasEdgeBetween(uid, vid string) bool {
	panic("not implemented")
}

func (g *donutGraph) NewNode(id string) Node {
	panic("not implemented")
}

func (g *donutGraph) NewEdge(uid, vid string, weight int) Edge {
	panic("not implemented")
}

func (g *donutGraph) Weight(uid, vid string) int {
	return 1
}

func (g *donutGraph) AddEdge(e Edge) {
	panic("not implemented")
}

func (g *donutGraph) GetEdge(from, to string) *Edge {
	panic("not implemented")
}

