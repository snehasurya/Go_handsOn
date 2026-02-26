package testing

import (
	"errors"
	"strings"
)

func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("The negative int")
	}
	return a + b, nil
}

func isPalindrom(input string) bool {
	s := strings.ToLower(input)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
