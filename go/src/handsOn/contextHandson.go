package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var i int
		for {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			select {
			case <-ctx.Done():
				return
			default:
				i++
			}
		}

	}()

	fmt.Println("context error", ctx.Err())
	time.Sleep(time.Second * 3)
	fmt.Println("no of go routines : ", runtime.NumGoroutine())
	cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("context error after cancel : ", ctx.Err())
	fmt.Println("no of go routines : ", runtime.NumGoroutine())
}
