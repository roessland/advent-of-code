package main

type Shortest struct {
	nodes []Node
	dist map[string]map[string]int
	absent int
}

func newShortest(nodes []Node) Shortest {
	as := Shortest{
		nodes: nodes,
		dist: make(map[string]map[string]int),
	}
	for _, node := range nodes {
		as.dist[node.ID()] = make(map[string]int)
	}
	return as
}

func (as Shortest) at(uid, vid string) int {
	return as.dist[uid][vid]
}

func (as Shortest) set(uid, vid string, weight int) {
	as.dist[uid][vid] = weight
}

func (as Shortest) Weight(uid, vid string) int {
	return as.dist[uid][vid]
}