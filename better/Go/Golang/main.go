package main

import "fmt"

func main() {
	input := "dqeurstaepiiaaueao"

	output := make(map[rune]int)
	output['a'] = 0
	output['e'] = 0
	output['i'] = 0
	output['o'] = 0
	output['u'] = 0

	for _, letter := range input {

		if _, ok := output[letter]; ok {
			output[letter]++
		}

	}
	for k, v := range output {
		fmt.Printf("%c : %d\n", k, v)
	}

}
