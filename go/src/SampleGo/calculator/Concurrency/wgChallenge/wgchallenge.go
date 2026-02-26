package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}

func printMsg() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello world"
	printMsg()
	wg.Add(1)
	go updateMsg("Hello Universe")
	wg.Wait()
	printMsg()

	wg.Add(1)
	go updateMsg("Hello cosmos")
	wg.Wait()
	printMsg()
}
