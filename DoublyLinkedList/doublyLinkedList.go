package main

import (
	"container/list"
	"fmt"
)

func main() {
	ll := list.New()
	ll.PushBack("a")
	ll.PushBack("b")
	ll.PushBack("c")
	ll.PushBack("d")
	ll.PushBack("e")
	ll.PushBack("f")

	for e := ll.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%s\t", e.Value)
	}
	fmt.Println()
	mark := ll.Front().Next()
	ll.InsertAfter("s", mark)
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %s -->", e.Value)
	}
	el := ll.Back()
	ll.Remove(el)
	fmt.Println()
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %s -->", e.Value)
	}

	for e := ll.Front(); e != nil; e = e.Next() {
		if e.Value == "c" {
			mark = e
			break
		}
	}
	if mark.Value != nil {
		ll.InsertAfter("z", mark)
	}
	fmt.Println()
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Printf(" %s -->", e.Value)
	}
}
