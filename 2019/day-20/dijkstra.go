package main

import (
	"container/heap"
	"math"
)

const infinity = math.MaxInt32

func Dijkstra(g Graph, sources []Node, isTarget func(Node) bool) (paths Shortest) {
	paths = NewShortest(infinity)

	for _, source := range sources {
		sid := source.ID()
		queue := NewPriorityQueue()

		// Seed priority queue with the source node
		paths.SetDist(sid, sid, 0)
		heap.Push(queue, PriorityHeapEntry{
			Priority: 0,
			Value:    sid,
		})

		for queue.Len() > 0 {
			// Pop node closest to source
			uID := heap.Pop(queue).(PriorityHeapEntry).Value
			u := g.Node(uID)

			// Early stopping if applicable
			if isTarget != nil && isTarget(u) {
				break
			}

			// Update neighbor nodes with closest distance
			suDist := paths.GetDist(sid, uID)
			for _, v := range g.From(uID) {
				vID := v.ID()
				svDist := paths.GetDist(sid, vID)
				uvDist := g.Weight(uID, vID)
				// Check if there is a shorter path to v
				if suDist+uvDist < svDist {
					// If updated, add to heap
					paths.SetDist(sid, vID, suDist+uvDist)
					heap.Push(queue, PriorityHeapEntry{
						Priority: suDist + uvDist,
						Value:    vID,
					})
				}
			}
		}
	}

	return paths
}
