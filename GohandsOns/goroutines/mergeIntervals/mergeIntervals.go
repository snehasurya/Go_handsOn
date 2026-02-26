package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3}, {8, 10}, {2, 6}, {15, 18}, {2, 6}}
	fmt.Println(intervals)
	ints := [][]int{
		{4, 7}, {1, 4}}

	fmt.Println(intervalsMerge(intervals))
	fmt.Println(intervalsMerge(ints))
}

func intervalsMerge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	var mergedIntervals [][]int
	mergedIntervals = append(mergedIntervals, intervals[0])
	for i := 0; i < len(intervals); i++ {
		nextInterval := intervals[i]
		lastMerged := mergedIntervals[len(mergedIntervals)-1]
		if nextInterval[0] <= lastMerged[1] {
			if nextInterval[1] >= lastMerged[1] {
				lastMerged[1] = nextInterval[1]
			}
		} else {
			mergedIntervals = append(mergedIntervals, nextInterval)
		}
	}
	return mergedIntervals
}
