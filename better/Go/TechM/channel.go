package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChan := make(chan string, 2) // Buffered channel with capacity 2

	go func() {
		fmt.Println("Sender: Sending 'one'...")
		bufferedChan <- "one" // Will not block immediately as capacity is 2
		fmt.Println("Sender: Sending 'two'...")
		bufferedChan <- "two" // Will not block immediately
		fmt.Println("Sender: Sending 'three'...")
		bufferedChan <- "three" // This will block until a value is received
		fmt.Println("Sender: 'three' sent.")
	}()

	time.Sleep(1 * time.Second) // Simulate some work before receiving
	fmt.Println("Receiver: Receiving first value...")
	msg1 := <-bufferedChan
	fmt.Printf("Receiver: Received '%s'\n", msg1)

	fmt.Println("Receiver: Receiving second value...")
	msg2 := <-bufferedChan
	fmt.Printf("Receiver: Received '%s'\n", msg2)

	fmt.Println("Receiver: Receiving third value...")
	msg3 := <-bufferedChan
	fmt.Printf("Receiver: Received '%s'\n", msg3)
}
