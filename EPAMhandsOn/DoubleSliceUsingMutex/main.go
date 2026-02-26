package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doubleSlice(input []int, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < end; i++ {
		input[i] = input[i] * 2
	}
}

func main() {
	//var wg sync.WaitGroup
	size := 100_000_000
	input := make([]int, size)

	//fmt.Println(input)
	for i := 0; i < size; i++ {
		input[i] = i + 1
	}

	// fmt.Printf("no of cores %d\n", runtime.NumCPU())
	// runtime.GOMAXPROCS(6)
	// fmt.Printf("no of cores %d\n", runtime.NumCPU())
	// fmt.Printf("maxpocs %d\n", runtime.GOMAXPROCS(0))
	// runtime.GOMAXPROCS(1)
	// fmt.Printf("maxpocs %d\n", runtime.GOMAXPROCS(0))
	// fmt.Printf("len of input %d\n", len(input))

	runtime.GOMAXPROCS(1)
	start := time.Now()
	runWork(input, 5)
	fmt.Printf("1 Core: %v\n", time.Since(start))
	runtime.GOMAXPROCS(runtime.NumCPU())
	start = time.Now()
	runWork(input, 5)
	fmt.Printf("all Core: %v\n", time.Since(start))

}

func runWork(input []int, workers int) {
	var wg sync.WaitGroup
	numOfGoroutines := workers
	chunks := len(input) / numOfGoroutines
	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		start := i * chunks
		end := start + chunks
		if i == numOfGoroutines-1 {
			end = len(input)
		}
		go doubleSlice(input, start, end, &wg)
	}
	wg.Wait()
}
