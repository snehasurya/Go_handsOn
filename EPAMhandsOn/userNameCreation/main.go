package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

const alphaNum = "abcdefghijklmnopqrstuvwxyz0123456789"

func sanitize(s string) string {
	var input string
	for _, letter := range []rune(s) {
		if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
			input += string(letter)
		}
	}
	return input
}
func segment(s string) string {
	for len(s) < 4 {
		s += string(alphaNum[rand.Intn(len(alphaNum))])
	}
	return s[:4]
}
func generateUserID(firstName, lastName string) string {
	return segment(sanitize(firstName)) + segment(sanitize(lastName))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateUserID("sn@h", "suryawanshi"))
}
