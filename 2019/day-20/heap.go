package main

import "container/heap"

type PriorityHeapEntry struct {
	Priority int
	Value string
}

type PriorityHeap []PriorityHeapEntry

func (h *PriorityHeap) Len() int {
	return len(*h)
}

func (h *PriorityHeap) Less(i, j int)bool {
	return (*h)[i].Priority < (*h)[j].Priority
}

func (h *PriorityHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PriorityHeap) Push(x interface{}) {
	entry := x.(PriorityHeapEntry)
	*h = append(*h, entry)
}