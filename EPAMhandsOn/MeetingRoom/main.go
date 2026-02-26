package main

import "fmt"

func findMeetingRoom(n int, meetings [][]int) int {
	start := make([]int, 0)
	end := make([]int, 0)

	for _, meet := range meetings {
		start = append(start, meet[0])
		end = append(end, meet[1])
	}
	room := 0
	startPrt := 0
	endPtr := 0
	for startPrt = 0; startPrt < n; startPrt++ {
		if start[startPrt] < end[endPtr] {
			room++
		} else {
			endPtr++
		}
	}
	return room
}

func main() {
	meetings := [][]int{
		{0, 5},
		{5, 10},
		{15, 20},
		{20, 30},
	}
	fmt.Println(findMeetingRoom(4, meetings))
	inp := 0
	fmt.Scanf("%d", inp)
	fmt.Println(inp)
}
