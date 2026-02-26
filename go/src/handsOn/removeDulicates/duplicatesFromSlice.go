// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func mainno() {
	input := []int{1, 2, 3, 4, 2, 3, 4}
	output, err := removeDuplicates(input)
	fmt.Println("output :  ", output)
	fmt.Println("error : ", err)

}

func removeDuplicates(input []int) ([]int, error) {
	if len(input) <= 1 {
		return input, errors.New("no suffitient element in slice")
	}
	var output []int
	map1 := make(map[int]bool)
	for _, in := range input {
		if _, ok := map1[in]; !ok {
			map1[in] = true
			output = append(output, in)
		}
	}

	return output, nil
}
