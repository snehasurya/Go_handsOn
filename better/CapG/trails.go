package main

import (
	"fmt"
	"math"
	"strings"
)

func mainnoo() {
	s := "1337c0d3"
	//s = strings.TrimSpace(s)
	//fmt.Println(math.MinInt32)
	fmt.Println(myAtoi(s))
	x := 1534236469
	fmt.Println(reverse(x))
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	fmt.Println(intToRoman(3749))
}

func reverse(x int) int {
	reverse := 0
	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	i, result, sign := 0, 0, 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		sign = 1
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if result > (math.MaxInt-digit)/10 {
			if sign == 1 {
				return math.MaxInt
			}
			return math.MinInt
		}
		result = result*10 + digit
		i++
	}
	return result * sign
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	output := make([]int, len(nums1)+len(nums2))
	for k := 0; k < (len(nums1) + len(nums2)); k++ {
		if i == len(nums1) {
			output[k] = nums2[j]
			j++
			continue
		}
		if j == len(nums2) {
			output[k] = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			output[k] = nums1[i]
			i++
		} else {
			output[k] = nums2[j]
			j++
		}
	}
	//fmt.Println(output)
	medianIndex := len(output) / 2
	if len(output)%2 == 0 {
		return (float64(output[medianIndex-1]) + float64(output[medianIndex])) / 2
	} else {
		return float64(output[medianIndex])
	}

}

func intToRoman(num int) string {
	value := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	symbol := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}
	var output string
	for num > 0 {
		for i := range value {
			if num >= value[i] {
				output += symbol[i]
				num -= value[i]
				break
			}
		}
	}
	return output
}
