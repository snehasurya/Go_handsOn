package main

import (
	"fmt"
	"time"
)

func printMsg(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printMsg("hello")
	printMsg("world")
	i := 1
	defer fmt.Println("hi")
	fmt.Println(i)
}
