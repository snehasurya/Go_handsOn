package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	ticker := time.NewTicker(700 * time.Millisecond)
	defer ticker.Stop()

	done := time.After(4 * time.Second)
	count := 0

	for {
		select {
		case <-done:
			{
				fmt.Println("time up")
				return
			}
		case t := <-ticker.C:
			{
				fmt.Printf("produced %d at %v\n", count, t.Format("05.000"))
				ch <- count
				count++
			}
		}
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("consumed %d\n", v)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
	fmt.Println("done")
}
