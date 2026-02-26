package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	chanIn := make(chan struct {
		index int
		value int
	})

	for i, v := range input {
		go func(i, v int) {
			chanIn <- struct {
				index int
				value int
			}{i, v * 2}
		}(i, v)
	}
	output := make([]int, len(input))
	//fmt.Println(output)
	for i := 0; i < len(input); i++ {
		res := <-chanIn
		output[res.index] = res.value
	}
	//fmt.Println(input)
	//fmt.Println(output)

	//ch := make(chan int)
	out := make(chan []int)
	go double(input, out)

	output1 := <-out
	fmt.Println(output1)
}

func double(input []int, output chan []int) {
	res := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		res[i] = input[i] * 3
	}
	output <- res
}
