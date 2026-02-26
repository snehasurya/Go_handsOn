package slice

import "fmt"

func main() {

	slice := []int{3, 2, 8, 2, 6, 5, 7}

	max := 0
	for _, val := range slice {
		if max < val {
			max = val
		}
	}
	fmt.Printf("max is %d \n", max)
	removeDuplicate(slice)
}

func removeDuplicate(input []int) {
	mapForValue := make(map[int]bool)
	output := []int{}
	for _, value := range input {
		if mapForValue[value] == false {
			mapForValue[value] = true
			output = append(output, value)
		}
	}
	fmt.Println(output)
}
