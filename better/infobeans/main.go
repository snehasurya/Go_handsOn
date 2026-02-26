/*** Problem Statement
You are given a list of integers.
Start N worker goroutines (N is configurable).
Each worker:
Reads numbers from a jobs channel
Computes the square of the number
Sends the result to a results channel
The program should:
Ensure all numbers are processed exactly once
Avoid goroutine leaks
Close channels correctly
Finally, collect and print all results.
Constraints / Expectations
Use goroutines and channels
Use sync.WaitGroup
No race conditions
Program should terminate cleanly ***/

package main

import (
	"fmt"
	"sync"
)

func worker(jobs, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for input := range jobs {
		result <- input * input
	}
}

func main() {
	jobs := make(chan int, 5)
	result := make(chan int, 5)
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numGoRoutines := 3
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go worker(jobs, result, &wg)
	}
	for _, job := range inputSlice {
		jobs <- job
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(result)
	}()
	var output []int
	for a := range result {
		output = append(output, a)
	}
	//wg.Wait()
	fmt.Println(output)

}
