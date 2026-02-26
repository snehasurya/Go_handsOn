package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) append(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	list.size++
}

func (list *LinkedList) printList() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil \n")
}
func (list *LinkedList) hasCycle() bool {
	if list.size < 2 {
		fmt.Println("no enough element in list to create cycle")
		return false
	} else {
		slow := list.head
		fast := list.head

		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				fmt.Println("There is a cycle")
				return true
			}
		}
	}
	return false
}

func (list *LinkedList) getNodeAt(index int) *Node {
	if index < 0 || index >= list.size {
		return nil
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current
}

func (list *LinkedList) createCycle(fromIndex, toIndex int) error {
	if list.size <= 2 {
		return fmt.Errorf("cannot create cycle no sufficient elements \n")
	}
	if fromIndex < 0 || fromIndex >= list.size || toIndex < 0 || toIndex >= list.size {
		return fmt.Errorf("invalid index \n")
	}
	fromNode := list.getNodeAt(fromIndex)
	toNode := list.getNodeAt(toIndex)
	if fromNode == nil || toNode == nil {
		return fmt.Errorf("unable to retrieve node at the index \n")
	}
	fromNode.next = toNode
	return nil
}
func main() {

	list1 := NewLinkedList()
	list1.append(10)
	list1.append(20)
	list1.append(30)
	list1.append(40)
	list1.printList()
	fmt.Println(list1.hasCycle())
	fmt.Println(list1.getNodeAt(2).value)
	// if err := list1.createCycle(3, 1); err != nil {
	// 	fmt.Printf("Error creating cycle: %v\n", err)
	// }
	// fmt.Printf("list has cycle : %t \n", list1.hasCycle())
	//list1.printList()
}
