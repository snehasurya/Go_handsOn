package main

import "fmt"

func producer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	//close(ch)
	return ch
}

func consumer(ch <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	// for c := range ch {
	// 	fmt.Println(c)
	// }
}

func main() {
	numChan := producer()
	consumer(numChan)
}
