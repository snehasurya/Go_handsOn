package main

import "fmt"

func main() {
	//input := []int{23, 2, 4, 6, 7}
	//fmt.Println(findSubArray(input, 13))
	//fmt.Println(findSumByMap(input, 12))
	//fmt.Println(findSumByMap(input, 29))
	slices()
	ch := make(chan int)
	go chan1(ch)
	for c := range ch {
		fmt.Println(c)
	}
	fmt.Printf("closed: %#v\n", <-ch)
	msg, ok := <-ch
	fmt.Printf("closed: %#v \t ok=%v\n", msg, ok)
	ch <- 1
}
func findSubArray(input []int, target int) []int {
	for i := 0; i < len(input); i++ {
		k := i
		sum := 0
		if input[i] > target {
			continue
		}
		for k <= len(input) {
			if sum < target {
				sum += input[k]
				k++
			} else if sum == target {
				return input[i:k]
			} else {
				break
			}
		}
	}
	return []int{}
}

// [23, 2, 4, 6, 7]
func findSumByMap(input []int, target int) []int {
	seen := make(map[int]int)
	seen[0] = -1
	currentSum := 0
	for i, v := range input {
		currentSum += v
		targetSum := currentSum - target
		if index, ok := seen[targetSum]; ok {
			return input[index+1 : i+1]
		}
		seen[currentSum] = i
	}
	return nil
}

func chan1(ch chan int) {
	//ch := make(chan int)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func slices() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	xp := [][]string{jb, jm}

	for i, v := range xp {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
