package main

import "fmt"

func calculate(input string) int {
	result := 0
	sign := 1
	num := 0
	//10-20
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		}
		if char == '+' || char == '-' || i == len(input)-1 {
			result += sign * num
			num = 0
			if char == '-' {
				sign = -1
			} else {
				sign = 1
			}
		}
	}
	return result
}
func main() {
	fmt.Println(calculate("10+20"))
	fmt.Println(calculate("20-10-4"))
}
