package main

import "fmt"

type Stack struct {
	items []int
}

func (st *Stack) Push(value int) {
	st.items = append(st.items, value)
}

func (st *Stack) Pop() (int, bool) {
	if st.IsEmpty() {
		return 0, false
	}
	lastIndex := len(st.items) - 1
	itemToPop := st.items[lastIndex]
	st.items = st.items[:lastIndex]
	return itemToPop, true
}

func (st *Stack) IsEmpty() bool {
	return len(st.items) == 0
}

func main() {
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println(stack.items)
	item, bools := stack.Pop()
	fmt.Println(item, " poped ", bools)
	fmt.Println(stack.items)

}
