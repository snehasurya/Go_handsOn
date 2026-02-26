package main

import (
	"fmt"
	"sort"
)

type friend struct {
	name  string
	grade string
}

type friends []friend

var grades = map[string]int{
	"A+": 1,
	"A":  2,
	"B+": 3,
	"B":  4,
	"C":  5,
}

func (f friends) Len() int {
	return len(f)
}
func (f friends) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f friends) Less(i, j int) bool {
	if grades[f[i].grade] > grades[f[j].grade] {
		return true
	}
	return false
}

func main() {
	frnd := []friend{
		{"sneha", "A"},
		{"vivaan", "A+"},
		{"pankaj", "B"},
		{"nanu", "C"},
	}
	sort.Sort(friends(frnd))
	fmt.Println(frnd)
}
