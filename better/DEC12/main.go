package main

import (
	"fmt"
	"time"
)

func ping(chan1 chan bool) {
	for i := 0; i < 5; i++ {
		<-chan1
		fmt.Println("Ping")
		chan1 <- true
	}

}

func pong(chan1 chan bool) {
	for i := 0; i < 5; i++ {
		<-chan1
		fmt.Println("Pong")
		chan1 <- true
	}
}

func main() {
	chan1 := make(chan bool)
	go ping(chan1)
	chan1 <- true
	go pong(chan1)
	time.Sleep(time.Second * 1)

}
