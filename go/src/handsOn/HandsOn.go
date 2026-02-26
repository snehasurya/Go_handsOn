package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func main() {
	input := []int{2, 3, 4, 5, 6, 7, 9, 11}
	var output []int
	for _, v := range input {
		if findPrimeNumber(v) {
			output = append(output, v)
		}
	}
	fmt.Println(output)
	ch := make(chan int, len(input))
	var output2 []int
	for _, v := range input {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			if findPrimeNumber(v) {
				ch <- v
			}
		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for n := range ch {
		output2 = append(output2, n)
	}
	fmt.Println(output2)
}

func findPrimeNumber(input int) bool {
	for i := 2; i <= int(math.Sqrt(float64(input))); i++ {
		if (input % i) == 0 {
			return false
		}
	}
	return true
}
