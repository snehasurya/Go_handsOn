package main

import (
	"fmt"
	"math"
)

func findPeaks(input []float64, threshold float64) int {
	peaks := make([]float64, 0)

	for i := 1; i < len(input)-1; i++ {
		isPeak := false
		diff := math.Abs(input[i] - input[i+1])
		diff2 := math.Abs(input[i] - input[i-1])
		//fmt.Println(diff)
		//fmt.Println(diff2)
		if threshold <= diff && threshold <= diff2 {
			isPeak = true
		}
		if isPeak {
			peaks = append(peaks, input[i])
		}
	}
	fmt.Println(peaks)
	return len(peaks)
}

func main() {
	input := []float64{8, 10.7, 17.1, 11.2, 13.5, 9.9, 14.9, 9.4, 9.4, 3.1, 12.7}
	threshold := 5.0
	fmt.Println(findPeaks(input, threshold))
}
