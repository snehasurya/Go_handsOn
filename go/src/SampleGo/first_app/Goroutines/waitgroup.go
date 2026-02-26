package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	input1 := []int{1, 2, 3, 4, 5}
	output1 := make([]int, len(input1))

	wg.Add(len(input1))
	for i, v := range input1 {
		go double(v, i, output1, &wg)
	}
	wg.Wait()
	fmt.Println(output1)
}

func double(input int, index int, output []int, wg *sync.WaitGroup) {
	defer wg.Done()
	output[index] = input * 2

}
