package main

import (
	"fmt"
	"sort"
)

type friend struct {
	name  string
	grade string
}

type byGrade []friend

var gradeOrder = map[string]int{
	"A+": 5,
	"A":  4,
	"B+": 3,
	"B":  2,
	"C":  1,
}

func (f byGrade) Len() int {
	return len(f)
}
func (f byGrade) Swap(a, b int) {
	f[a], f[b] = f[b], f[a]
}
func (f byGrade) Less(a, b int) bool {
	gradeA := gradeOrder[f[a].grade]
	gradeB := gradeOrder[f[b].grade]
	return gradeA > gradeB
}
func main() {
	friends := []friend{
		{name: "Aima", grade: "A"},
		{name: "Arjun", grade: "A+"},
		{name: "Iram", grade: "B+"},
		{name: "bram", grade: "C"},
		{name: "cram", grade: "B"},
	}
	sort.Sort(byGrade(friends))
	fmt.Println(friends)
}
