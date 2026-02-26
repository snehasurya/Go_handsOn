package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	input := make([]int, 0)
	output := make([]int, 0)

	for i := 0; i < 100; i++ {
		input = append(input, i+1)
	}
	job := make(chan int, 7)
	result := make(chan int, 7)

	numGoRoutines := 5

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go doubleTheSlice(job, result, &wg)
	}

	go func() {
		for _, task := range input {
			job <- task
		}
		close(job)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for values := range result {
		output = append(output, values)
	}
	fmt.Println(output)
}

func doubleTheSlice(job, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range job {
		result <- num * 2
	}
}
