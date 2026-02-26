package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Odd length palindromes (centered at s[i])
		len1 := expandAroundCenter(s, i, i)
		// Even length palindromes (centered between s[i] and s[i+1])
		len2 := expandAroundCenter(s, i, i+1)

		maxLength := max(len1, len2)

		if maxLength > end-start {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s1 := "babad"
	fmt.Printf("Longest palindromic substring in \"%s\": \"%s\"\n", s1, longestPalindrome(s1)) // Output: "bab" or "aba"

	s2 := "cbbd"
	fmt.Printf("Longest palindromic substring in \"%s\": \"%s\"\n", s2, longestPalindrome(s2)) // Output: "bb"

	s3 := "a"
	fmt.Printf("Longest palindromic substring in \"%s\": \"%s\"\n", s3, longestPalindrome(s3)) // Output: "a"

	s4 := "ac"
	fmt.Printf("Longest palindromic substring in \"%s\": \"%s\"\n", s4, longestPalindrome(s4)) // Output: "a" or "c"
}
