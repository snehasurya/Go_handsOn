package main

import "fmt"

func main() {

	//chan1 := make(chan bool)
	input := []int{1, 2, 3, 4, 5}
	output := make([]int, len(input))
	chan2 := make(chan bool, len(input))
	// go func() {
	// 	for i := range input {
	// 		input[i] = input[i] * 2
	// 	}
	// 	chan1 <- true
	// }()
	// <-chan1
	// fmt.Println(input)
	for i, v := range input {
		go doubleSlice(v, output, i, chan2)
	}
	for i := 0; i < len(input); i++ {
		<-chan2
	}
	fmt.Println(output)
}

func doubleSlice(input int, output []int, index int, chan1 chan bool) {
	output[index] = input * 2
	chan1 <- true
}
