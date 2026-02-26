package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func tricky(result int) int {
	defer func() {
		result++
		fmt.Println(result)
	}()

	return result
}

func world() {
	defer fmt.Println("world")     // Deferred call 2
	defer fmt.Println("beautiful") // Deferred call 1
	fmt.Println("hello")
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}
func main() {
	fmt.Println(tricky(1))
	world()
	fmt.Println(c())
}
