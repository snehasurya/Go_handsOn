package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Printf("nil\n")
	fmt.Println("list size is: ", ll.size)
}

func (ll *LinkedList) append(value int) {
	newNode := &Node{
		value: value,
	}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}
func (ll *LinkedList) deleteNode(value int) bool {
	current := ll.head
	var prev *Node
	if ll.head.value == value {
		temp := ll.head
		ll.head = ll.head.next
		temp.next = nil
	}
	for current.next != nil {
		if current.value == value {
			temp := current
			prev.next = current.next
			temp.next = nil
		} else {
			prev = current
			current = current.next
		}
	}
	ll.size--
	return true
}
func (ll *LinkedList) appendAtIndex(value int, index int) {
	if index < 0 || index > ll.size {
		fmt.Errorf("index out of bound")
	}
	current := ll.head
	newNode := &Node{
		value: value,
	}
	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode
	} else {
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	ll.size++
}

func (ll *LinkedList) getIndexAt(index int) *Node {
	if index < 0 || index > ll.size {
		fmt.Errorf("index out of bound")
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}
func (ll *LinkedList) createCycle(fromIndex, toIndex int) error {
	if ll.size <= 2 {
		return errors.New("Not enough elements to create cycle")
	}
	if fromIndex < 0 || fromIndex > ll.size || toIndex < 0 || toIndex > ll.size {
		return errors.New("index out of bound")
	}
	fromIndexNode := ll.getIndexAt(fromIndex)
	toIndexNode := ll.getIndexAt(toIndex)
	if fromIndexNode == nil || toIndexNode == nil {
		return errors.New("no element at index")
	}
	fromIndexNode.next = toIndexNode
	return nil
}
func (ll *LinkedList) hasCycle() bool {
	fast := ll.head
	slow := ll.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}
func main() {

	list := NewLinkedList()
	list.append(10)
	list.append(20)
	list.append(30)
	list.printList()
	//list.deleteNode(10)
	//list.printList()
	list.appendAtIndex(40, 3)
	list.printList()
	//list.createCycle(3, 2)
	//list.printList()
	fmt.Println("list has a cycle : ", list.hasCycle())
}
