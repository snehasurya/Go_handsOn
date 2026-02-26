package main

import (
	"fmt"
	"time"
)

func odd(input chan int) {
	for i := 0; i < 5; i++ {
		num := <-input
		fmt.Println(num)
		input <- num + 1
	}
}

func even(input chan int) {
	for i := 0; i < 5; i++ {
		num := <-input
		fmt.Println(num)
		input <- num + 1
	}
}

func main() {
	input := make(chan int)
	go odd(input)
	go even(input)
	input <- 1
	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
