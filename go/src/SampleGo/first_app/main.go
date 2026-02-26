package main

import (
	"fmt"
	"strings"
)

func main() {
	const age = 30
	fmt.Println("age is  ", age)

	makeMap()
	fmt.Println("Hi Sneha")
	freq := findFrequesncy("Hey I am Sneha Sneha Hi")
	for word, fre := range freq {
		fmt.Printf("%s: %d  \n", word, fre)
	}
	fmt.Printf("factorial of %d is  : %d", 5, factorial(5))
	fmt.Printf("\nfibonacci  of %d is  : %d \n", 10, fibonacci(10))

	fmt.Println("Fibonacci sequence in recursion : ")
	for i := range 10 {
		fmt.Print(fibn(i), " ")
	}
	inputS := []int{1, 2, 3, 4, 5, 6}
	reverseInPlace(inputS)
	fmt.Println(inputS)
}

func findFrequesncy(input string) map[string]int {

	mapWithCount := make(map[string]int)
	stringSplit := strings.Fields(input)
	split := strings.Split(input, " ")

	for _, s := range stringSplit {
		mapWithCount[s]++
	}
	fmt.Println(split)
	return mapWithCount
}

func factorial(inputValue int) int64 {
	fact := 1
	input := inputValue
	for input > 1 {
		fact = fact * input
		input--
	}
	return int64(fact)
}

func fibonacci(input int) []int {
	if input <= 0 {
		return []int{}
	}
	seq := make([]int, input)
	seq[0] = 0
	if input >= 1 {
		seq[1] = 1
	}
	for i := 2; i < input; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq

}

func fibn(input int) int {
	if input <= 1 {
		return 1
	}
	return fibn(input-1) + fibn(input-2)
}

func reverseInPlace(input []int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func makeMap() {
	course := make(map[string]float64, 2)

	course["go"] = 4.5
	course["java"] = 4.8
	course["mainf"] = 4.3
	fmt.Println(course)
}
