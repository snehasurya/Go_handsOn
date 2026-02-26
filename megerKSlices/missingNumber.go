package main

import "fmt"

func missingNumber(arr []int) int {
	xor := 0
	n := len(arr)
	for i := 0; i <= n; i++ {
		xor ^= i
	}
	for _, v := range arr {
		xor ^= v
	}
	return xor
}

func main() {
	arr := []int{0, 1, 3}
	fmt.Println("Input:", arr)
	fmt.Println("Missing number:", missingNumber(arr))
}
