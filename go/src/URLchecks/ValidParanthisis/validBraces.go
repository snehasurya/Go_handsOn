package main

import "fmt"

func main() {
	fmt.Println(isValid(")()"))
}
func isValid(s string) bool {
	braces := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []rune
	for _, v := range s {
		if _, ok := braces[v]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || braces[stack[len(stack)-1]] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
