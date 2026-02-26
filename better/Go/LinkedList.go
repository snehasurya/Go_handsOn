package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) append(value int) {
	newNode := &Node{value: value}
	if ll.size == 0 {
		ll.head = newNode
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	ll.size++
}
func (ll *LinkedList) printList() {
	fmt.Println("size of list : ", ll.size)
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("nil\n")
}

func (ll *LinkedList) delete(input int) bool {
	if ll.head.value == input {
		temp := ll.head
		ll.head = temp.next
		temp.next = nil
		ll.size--
		return true
	} else {
		var prevNode *Node
		currentNode := ll.head
		for currentNode.next != nil {
			if currentNode.value == input {
				temp := currentNode
				prevNode.next = currentNode.next
				temp.next = nil
				ll.size--
				return true
			} else {
				prevNode = currentNode
				currentNode = currentNode.next
			}

		}
	}
	return false
}
func main() {
	list := newLinkedList()
	list.append(10)
	list.append(20)
	list.append(30)
	list.append(40)
	list.printList()
	fmt.Println(list.delete(50))

	list.printList()
}
