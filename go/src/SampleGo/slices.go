package main

import (
	"fmt"
)

func slices() {
	sl := make([]int, 3, 3)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	fmt.Printf("%d   %d", len(sl), cap(sl))
	sl = append(sl, 4)
	fmt.Println()
	fmt.Printf("%d   %d", len(sl), cap(sl))
	sl = append(sl, 4)
	sl = append(sl, 5)
	fmt.Println()
	fmt.Printf("%d   %d", len(sl), cap(sl))
	sl = append(sl, 6)
	fmt.Println()
	fmt.Printf("%d   %d", len(sl), cap(sl))
}

type product struct {
	name  string
	color string
}

const (
	red int = iota
	green
	blue
)

func structAndPointer() {
	p := []product{
		{"apple", "red"},
		{"orange", "orange"},
		{"banana", "yellow"},
	}
	for i, v := range p {
		fmt.Println(&v)
		fmt.Println(&p[i])
	}
}
func main() {
	fmt.Println(red)
	fmt.Println(blue)
	fmt.Println(green)

}
