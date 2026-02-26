package main

import "fmt"

func printEvenCha(evenChan, oddChan chan int, done chan bool) {
	for {
		num, ok := <-evenChan
		if !ok {
			done <- true
			return
		}
		fmt.Println("Even go routine:  n", num)
		oddChan <- num + 1
	}
}
func printOddCha(oddChan, evenChan chan int, done chan bool) {
	for {
		num, ok := <-oddChan
		if !ok {
			done <- true
			return
		}
		fmt.Println("Odd go routine:  ", num)
		if num+1 > 10 {
			close(evenChan)
			close(oddChan)
			return
		}
		evenChan <- num + 1
	}
}
func main() {

	evenChan := make(chan int)
	oddChan := make(chan int)
	done := make(chan bool)

	go printEvenCha(evenChan, oddChan, done)
	go printOddCha(evenChan, oddChan, done)
	evenChan <- 0

	<-done
	<-done

	fmt.Println("Done printing all routines")
}
