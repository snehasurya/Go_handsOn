package main

import (
	"fmt"
	"sync"
)

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, 5)
	result := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doubleSlice(ch, result, &wg)
	}
	for i := 0; i < len(input); i++ {
		ch <- input[i]
	}
	close(ch)
	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println(v)
	}
}

func doubleSlice(ch, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		result <- num * 2
	}
}
