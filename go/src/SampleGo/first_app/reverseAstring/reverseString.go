package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "Nitin"
	reverseString := reverse(inputString)
	fmt.Println("Input String", inputString)
	fmt.Println("reverseString", reverseString)
	isPalindromString := isPalindrom(inputString, reverseString)
	fmt.Println(isPalindromString)
}

func reverse(input string) string {
	input = strings.ToLower(input)
	runes := []rune(input)
	fmt.Println(runes)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrom(input, reverse string) bool {
	input = strings.ToLower(input)
	return input == reverse
}
