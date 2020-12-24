package main

type AllShortest struct {
	nodes []Node
	dist map[string]map[string]int
	absent int
}

func newAllShortest(nodes []Node) AllShortest {
	as := AllShortest{
		nodes: nodes,
		dist: make(map[string]map[string]int),
	}
	for _, node := range nodes {
		as.dist[node.ID()] = make(map[string]int)
	}
	return as
}

func (as AllShortest) at(uid, vid string) int {
	return as.dist[uid][vid]
}

func (as AllShortest) set(uid, vid string, weight int) {
	as.dist[uid][vid] = weight
}

func (as AllShortest) Weight(uid, vid string) int {
	return as.dist[uid][vid]
}