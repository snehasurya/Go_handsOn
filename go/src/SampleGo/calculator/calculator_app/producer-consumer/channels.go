package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("producer ", i)
		ch <- i
	}

}

func consumer(cha <-chan int) {
	for i := range cha {
		fmt.Println(i)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)

	bufChan := make(chan int, 4)

	bufChan <- 10
	bufChan <- 11
	bufChan <- 12
	go func() {
		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)
	}()
	time.Sleep(2 * time.Second)
}
