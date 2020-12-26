package main

import "fmt"

func FloydWarshall(g Graph) (paths Shortest) {
	paths = NewShortest(infinity)
	nodes := g.Nodes()
	for _, u := range g.Nodes() {
		for _, v := range g.Nodes() {
			paths.SetDist(u.ID(), v.ID(), g.Weight(u.ID(), v.ID()))
		}
	}

	for _, k := range nodes {
		fmt.Println(k, len(nodes))
		kid := k.ID()
		for _, u := range nodes {
			uid := u.ID()
			for _, v := range nodes {
				vid := v.ID()
				alt := paths.GetDist(uid, kid) + paths.GetDist(kid, vid)
				if paths.GetDist(uid, vid) > alt {
					paths.SetDist(uid, vid, alt)
				}
			}
		}
	}
	return paths
}