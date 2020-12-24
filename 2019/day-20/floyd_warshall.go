package main

import "fmt"

func FloydWarshall(g Graph) (paths AllShortest) {
	paths = newAllShortest(g.Nodes())
	nodes := paths.nodes
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