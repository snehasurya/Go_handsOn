package main

import (
	"fmt"
	"sync"
)

func printTheInput(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	input := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
	}
	wg.Add(4) // if 5 it gives deadlock
	for i, x := range input {
		go printTheInput(fmt.Sprintf("%d : %s ", i, x), &wg)
	}
	//time.Sleep(1 * time.Nanosecond)
	wg.Wait()
	wg.Add(1) // if not adding this, gives panic: sync: negative WaitGroup counter
	printTheInput("second", &wg)
}
