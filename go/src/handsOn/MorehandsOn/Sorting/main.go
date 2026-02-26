package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (b ByAge) Len() int      { return len(b) }
func (b ByAge) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByAge) Less(i, j int) bool {
	if b[i] != b[j] {
		return b[i].Age < b[j].Age
	}
	return b[i].Name < b[j].Name
}

func main() {
	slice1 := []int{3, 5, 1, 2, 6}

	//sort.Ints(slice1)
	//fmt.Println(slice1)
	intSlice := sort.IntSlice(slice1)
	fmt.Println(intSlice)
	fmt.Println(sort.Reverse(intSlice))
	//sort.Sort(sort.Reverse(intSlice))
	//fmt.Println(slice1)

	people := []Person{
		{"Sneha", 30},
		{"Pankaj", 32},
		{"Vivaan", 3},
		{"Nanu", 3},
	}
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
