package main

import (
	"fmt"
	"sort"
)

func mainno() {
	input := []int{2, 1, 5, 1, 3, 2}
	size := 3
	subArray, sum := maxSumArray(input, size)
	fmt.Printf("SubArray : %v  and sum : %d  \n", subArray, sum)

}

func findMaxSum(input []int, size int) ([]int, int) {
	sort.Ints(input)
	sum := 0
	result := make([]int, 0, size)
	for i := 1; i <= size; i++ {
		result = append(result, input[len(input)-i])
		sum += input[len(input)-i]
	}
	return result, sum
}

func maxSumArray(input []int, size int) ([]int, int) {
	max := 0
	var result []int
	for i := 0; i <= len(input)-size; i++ {
		sum := 0
		subArray := input[i : i+size]
		for _, v := range subArray {
			sum += v
		}
		if max < sum {
			max = sum
			result = subArray
		}
	}
	return result, max
}
