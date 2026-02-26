package main

import (
	"fmt"
	"sync"
)

var name string = "abc"
var wg sync.WaitGroup

//var mu sync.Mutex

func main() {
	chan1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go accessor()
		chan1 <- i
		go producer(chan1)
		wg.Wait()
	}

}

func accessor() {
	defer wg.Done()
	fmt.Println(name)
}

func producer(count chan int) {
	defer wg.Done()
	name = fmt.Sprintf("%s%d", "qwe", <-count)

}
