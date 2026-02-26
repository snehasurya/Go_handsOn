package main

import (
	"fmt"
)

func main() {
	// An unbuffered channel with a capacity of zero

	var i interface{} = "sneha"
	fmt.Printf("%t", i)
	s := i.(string)
	fmt.Println(s)
	checktype(i)
	sl := new([]int)
	*sl = append(*sl, 1)
	*sl = append(*sl, 1)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(*sl), cap(*sl), sl)
	*sl = append(*sl, 1)
	*sl = append(*sl, 1)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(*sl), cap(*sl), sl)
	*sl = append(*sl, 1)

	fmt.Println(*sl)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(*sl), cap(*sl), sl)

	sl2 := make([]int, 0, 2)
	sl2 = append(sl2, 1)
	sl2 = append(sl2, 2)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(sl2), cap(sl2), sl2)
	sl2 = append(sl2, 3)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(sl2), cap(sl2), sl2)
	sl2 = append(sl2, 4)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(sl2), cap(sl2), sl2)
	sl2 = append(sl2, 4)
	fmt.Printf("After 4th append: len=%d, cap=%d, slice=%v\n", len(sl2), cap(sl2), sl2)
}

func checktype(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("it's int")
	case string:
		fmt.Println("it's string")
	}

}
