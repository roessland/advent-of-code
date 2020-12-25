package main

import (
	"fmt"
	"math"
)

const infinity = math.MaxInt32

func Dijkstra(g Graph, from []Node) (paths Shortest) {
	nodes := g.Nodes()
	paths = newShortest(nodes)

	for _, n0 := range from {
		var Q := NewSet()
		for _, u := range nodes {
			paths.set(n0.ID(), u.ID(), infinity)
		}
		paths.set(n0.ID(), n0.ID(), 0)
	}



	for _, u := range g.Nodes() {
		for _, v := range g.Nodes() {
			paths.set(u.ID(), v.ID(), g.Weight(u.ID(), v.ID()))
		}
	}

	for _, k := range nodes {
		fmt.Println(k, len(nodes))
		kid := k.ID()
		for _, u := range nodes {
			uid := u.ID()
			for _, v := range nodes {
				vid := v.ID()
				alt := paths.at(uid, kid) + paths.at(kid, vid)
				if paths.at(uid, vid) > alt {
					paths.set(uid, vid, alt)
				}
			}
		}
	}
	return paths
}