package main

import "fmt"

func main() {
	input := []int{0, 1, 0, 1, 0, 1, 0, 1, 0}
	fmt.Println(arrange(input))

}
func arrange(input []int) []int {
	left := 0
	right := len(input) - 1

	for left < right {
		for left < right && input[left] == 0 {
			left++
		}
		for left < right && input[right] == 1 {
			right--
		}

		if left < right {
			input[left] = 0
			input[right] = 1
			left++
			right--
		}
	}
	fmt.Println(left, right)
	return input
}

func swaps(input []int) []int {
	nextZero := 0
	//{0, 1, 0, 1, 0, 1, 0, 1, 0}
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			input[i], input[nextZero] = input[nextZero], input[i]
			nextZero++
		}
	}
	return input
}
