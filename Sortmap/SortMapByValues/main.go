package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key   string
	value int
}

type listOfPairs []pair

func (l listOfPairs) Len() int {
	return len(l)
}
func (l listOfPairs) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// Ascending sort by Value: element 'i' is less than element 'j'
func (l listOfPairs) Less(i, j int) bool { return l[i].value < l[j].value }

func main() {
	myMap := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
		"date":   4,
	}
	list1 := make(listOfPairs, len(myMap))
	i := 0
	for k, v := range myMap {
		list1[i] = pair{k, v}
		i++
	}
	fmt.Println(list1)

	//sort.Sort(list1)
	//fmt.Println(list1)
	sliceOflists := make([]pair, len(myMap))
	j := 0
	for k, v := range myMap {
		sliceOflists[j] = pair{k, v}
		j++
	}
	fmt.Println(sliceOflists)
	sort.Slice(sliceOflists, func(i, j int) bool {
		return sliceOflists[i].value < sliceOflists[j].value
	})
	fmt.Println(sliceOflists)
}
