package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := &minHeap{}
	heap.Init(h)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
		fmt.Println(*list)
	}
	newList := &ListNode{}
	current := newList
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return newList.Next
}

func sliceToList(input []int) *ListNode {
	if len(input) == 0 {
		fmt.Println("no sufficient elements in input")
		return nil
	}
	head := &ListNode{Val: input[0]}
	current := head
	for i := 1; i < len(input); i++ {
		current.Next = &ListNode{Val: input[i]}
		current = current.Next
	}
	return head
}
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Printf(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}
func main() {
	list1 := sliceToList([]int{1, 4, 5})
	list2 := sliceToList([]int{3, 6, 9})
	list3 := sliceToList([]int{2, 7, 8})
	printList(list1)
	printList(list2)
	printList(list3)
	lists := []*ListNode{list1, list2, list3}
	newList := mergeKLists(lists)
	printList(newList)

}
