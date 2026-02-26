package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Chmod(0777)
	b := []byte("sneha")
	f.Write(b)
	var c []byte
	c, _ = os.ReadFile("output.txt")
	fmt.Println(string(c))

}
