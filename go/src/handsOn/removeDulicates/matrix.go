package main

import "fmt"

func main() {
	slice1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := range slice1 {
		fmt.Printf("%d   ", slice1[i][0])
	}

}
