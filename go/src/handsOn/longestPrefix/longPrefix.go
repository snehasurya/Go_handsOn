package main

import (
	"fmt"
)

func main() {
	//strs := []string{"fllower", "fllow", "fllight"}
	strs1 := []string{"dog", "racecar", "car"}
	//fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix(strs1))
	isValid("{}()")
}
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = findPrefix(prefix, strs[i])
	}
	if prefix == "" {
		return ""
	}
	return prefix
}
func findPrefix(s1, s2 string) string {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}
	for i := 0; i < minLength; i++ {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:minLength]
}

func isValid(s string) bool {
    // 1. Initial setup (Stack and map)
    stack := make([]rune, 0)
    pairs := map[rune]rune{ ... } // The matching rules map

    // 2. Loop through every character in the string 's'
    for _, char := range s {

        // --- PART A: IS IT AN OPENING BRACKET? ---
        if _, isClosing := pairs[char]; !isClosing {
            // Yes, it is an opening bracket (like '(', '{', or '[').
            // Push it onto the stack and continue to the next character.
            stack = append(stack, char)
            continue
        }

        // --- PART B: IS IT A CLOSING BRACKET? ---

        // Check 1: If the stack is empty, but we see a CLOSING bracket (like ']').
        // This is invalid: "}"
        if len(stack) == 0 {
            return false 
        }

        // Check 2: Does the CLOSING bracket match the last OPENING one?
        // Get the required opener (e.g., if char is ')', requiredOpen is '(')
        requiredOpen := pairs[char] 
        
        // Get the actual last opener from the top of the stack
        lastOpen := stack[len(stack)-1] 

        if lastOpen == requiredOpen {
            // They match! Remove the last opener from the stack (POP).
            stack = stack[:len(stack)-1]
        } else {
            // They DON'T match (e.g., top is '(' but char is ']'). Invalid!
            return false
        }
    }

    // 3. Final Check: Did every opener get closed?
    // If the stack is empty, every bracket had a match. Valid!
    // If the stack still has items, some openers were never closed. Invalid!
    return len(stack) == 0
}