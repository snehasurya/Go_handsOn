package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{2, 4, 5, 6, 7, 8}
	var output []int

	for _, i := range input {
		if primeNumer(i) {
			output = append(output, i)
		}
	}
	fmt.Print(output)
}

func primeNumer(input int) bool {
	for k := 2; k <= int(math.Sqrt(float64(input))); k++ {
		if (input % k) == 0 {
			return false
		}
	}
	return true
}
