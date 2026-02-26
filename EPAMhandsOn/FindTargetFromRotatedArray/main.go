package main

import "fmt"

//0  1  2  3  4  5  6
//{4, 5, 6, 7, 8, 9, 1, 2, 3}

func findTarget(input []int, target int) int {
	left := 0
	right := len(input) - 1

	for left <= right {
		mid := left + (right-left)/2

		if target == input[mid] {
			return mid
		}
		if input[left] <= input[mid] {
			if input[left] <= target && target < input[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if input[mid] < target && target <= input[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findTarget([]int{4, 5, 6, 7, 1, 2, 3}, 6))
}
