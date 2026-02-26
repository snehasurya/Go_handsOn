package main

import "fmt"

func getletter(n int) string {
	return string(rune('a' + n - 1))
}

func decode(digits string, current string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, current)
		return
	}
	if digits[0] != '0' {
		val := int(digits[0] - '0')
		decode(digits[1:], current+getletter(val), result)
	}
	if len(digits) >= 2 {
		val := int(digits[0]-'0')*10 + int(digits[1]-'0')
		if val >= 10 || val <= 26 {
			decode(digits[2:], current+getletter(val), result)
		}

	}
}
func main() {
	input := "1123"
	result := make([]string, 0)
	decode(input, "", &result)
	for _, v := range result {
		fmt.Println(v)
	}
}
