package main

import "fmt"

func longestSubString(input string) int {
	charMap := make(map[rune]int)
	left := 0
	maxLength := 0

	for right, char := range []rune(input) {

		if lastIndex, found := charMap[char]; found && lastIndex >= left {
			left = lastIndex + 1
		}
		charMap[char] = right
		currentLength := right - left + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func main() {
	input := "abcdabcdeabccd"
	fmt.Println(longestSubString(input))
}
