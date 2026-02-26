package main

import (
	"fmt"
	"time"
)

func main() {
	cha1 := make(chan int)
	cha2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		cha1 <- 10
	}()
	go func() {
		cha2 <- 20
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-cha1:
			fmt.Println(msg1)

		case msg2 := <-cha2:
			fmt.Println(msg2)
		}
	}

}
