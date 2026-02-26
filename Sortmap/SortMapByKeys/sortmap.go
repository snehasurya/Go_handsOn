package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. The original unordered map
	myMap := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
		"date":   4,
	}

	// 2. Create a slice to hold the keys
	keys := make([]string, 0, len(myMap))
	for key := range myMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	sortedMap := make(map[string]int, len(myMap))
	for _, key := range keys {
		sortedMap[key] = myMap[key]
		//fmt.Printf("%s : %d \n", key, myMap[key])
	}
	fmt.Println(sortedMap)
}
