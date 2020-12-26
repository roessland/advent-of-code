package main

import "container/heap"

type PriorityHeapEntry struct {
	Priority int
	Value string
}

type PriorityQueue []PriorityHeapEntry

var _ heap.Interface = &PriorityQueue{}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (h *PriorityQueue) Len() int {
	return len(*h)
}

func (h *PriorityQueue) Less(i, j int)bool {
	return (*h)[i].Priority < (*h)[j].Priority
}

func (h *PriorityQueue) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	entry := x.(PriorityHeapEntry)
	*h = append(*h, entry)
}

func (h *PriorityQueue) Pop() interface{} {
	entry := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return entry
}