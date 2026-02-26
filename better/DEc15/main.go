package main

import "fmt"

func main() {

	input := "((My Name ( is ( Sne ) ha ) su)ryawanshi)"
	fmt.Println(findMatchedBraces(input, 0))
}

func findMatchedBraces(input string, position int) int {
	stack := make([]int, 0, 0)
	type pair struct {
		open  int
		close int
	}
	var match []pair
	for i, v := range input {
		if v == '(' {
			stack = append(stack, i)
		} else if v == ')' {
			openIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			match = append(match, pair{open: openIndex, close: i})
		}
	}
	for _, m := range match {
		if m.open == position {
			return m.close
		}
		fmt.Printf("%d  %d \n", m.open, m.close)
	}
	return 0
}
