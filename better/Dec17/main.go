//str := "aaaaababbbaaaccccabcdaaabaccdddd"
//searchStr := "aba"

package main

import "fmt"

type input struct {
	str string
}

type findString interface {
	findSubString(string) int
}

func (in input) findSubString(target string) int {
	for i := 0; i < len(in.str)-len(target); i++ {
		if in.str[i:i+len(target)] == target {
			return i
		} else {
			continue
		}
	}
	return 0
}

func main() {
	in := input{
		str: "aaaaababbbaaaccccabcdaaabaccdddd",
	}
	fmt.Println(in.findSubString("aba"))
	var fd findString
	fd = in
	fmt.Println(fd.findSubString("aba"))
}
