package main

import (
	"container/heap"
	"fmt"
)

type element struct {
	value    int
	seqId    int
	eleIndex int
}

type minHeap []element

func (h minHeap) Len() int      { return len(h) }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h minHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	return x
}

func mergeKSlices(arr [][]int) []int {
	h := &minHeap{}
	heap.Init(h)
	for i := range arr {
		if len(arr[i]) > 0 {
			heap.Push(h, element{
				arr[i][0], i, 0})
		}
	}
	result := []int{}
	for h.Len() > 0 {
		ele := heap.Pop(h).(element)
		result = append(result, ele.value)
		nextInd := ele.eleIndex + 1
		if nextInd < len(arr[ele.seqId]) {
			heap.Push(h, element{arr[ele.seqId][nextInd], ele.seqId, nextInd})
		}
	}
	return result
}

func mainno() {
	input := [][]int{
		{1, 3, 4},
		{2, 5, 7},
		{6, 8, 9},
	}
	result := mergeKSlices(input)
	fmt.Println(result)
}
