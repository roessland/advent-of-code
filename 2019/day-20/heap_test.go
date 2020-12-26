package main

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHeap(t *testing.T) {
	queue := NewPriorityQueue()
	require.Equal(t, 0, queue.Len())

	heap.Push(queue, PriorityHeapEntry{
		Priority: 1,
		Value:    "one",
	})
	require.Equal(t, 1, queue.Len())


	heap.Push(queue, PriorityHeapEntry{
		Priority: 0,
		Value:    "zero",
	})
	require.Equal(t, 2, queue.Len())

	heap.Push(queue, PriorityHeapEntry{
		Priority: 2,
		Value:    "two",
	})
	require.Equal(t, 3, queue.Len())


	fmt.Println(heap.Pop(queue).(PriorityHeapEntry).Value)
	require.Equal(t, 2, queue.Len())

	fmt.Println(heap.Pop(queue).(PriorityHeapEntry).Value)
	require.Equal(t, 1, queue.Len())

	fmt.Println(heap.Pop(queue).(PriorityHeapEntry).Value)
	require.Equal(t, 0, queue.Len())


}
