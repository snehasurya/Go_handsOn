package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Ball struct{}

func player(name string, winner chan<- string, in <-chan Ball, out chan<- Ball, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			winner <- name
			return
		case ball := <-in:
			delay := time.Duration(rand.Intn(11)) * time.Second
			time.Sleep(delay)
			fmt.Printf("%s hits the ball after %v\n", name, delay)

			select {
			case out <- ball:
			case <-ctx.Done():
				winner <- name
				return
			}
		}
	}
}

func main() {
	//rand.Seed(1011)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	ping := make(chan Ball)
	pong := make(chan Ball)
	winner := make(chan string)
	go player("player A", winner, ping, pong, ctx)
	go player("player B", winner, pong, ping, ctx)
	fmt.Println("goroutines started")

	ping <- Ball{}

	w := <-winner
	fmt.Printf("winner is %s\n", w)
}
