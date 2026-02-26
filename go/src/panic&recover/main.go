package main

import "fmt"

func main() {
	fmt.Println(maths(20, 3))
	fmt.Println(maths(20, 0))
	fmt.Println(maths(20, 2))
}
func maths(inp1, inp2 int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered")
		}
	}()
	if inp2 == 0 {
		panic("inp2 is 0")
	}
	return inp1 / inp2
}
