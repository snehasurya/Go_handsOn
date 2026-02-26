package main

import (
	"fmt"
	"sync"
)

var count = 1
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 1)
	fmt.Println("hello")
	ch <- "sneha"
	fmt.Println("sent")
	msg := <-ch
	fmt.Println(msg)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go counter(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func counter(wg *sync.WaitGroup) {
	defer wg.Done()
	//mu.Lock()
	count++
	//mu.Unlock()
}
