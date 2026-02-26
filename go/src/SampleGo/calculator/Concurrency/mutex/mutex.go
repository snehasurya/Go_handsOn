package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}
func main() {
	var mutex sync.Mutex

	msg = "hey"
	wg.Add(2)
	go updateMsg("Hello world", &mutex)
	go updateMsg("Hello Universe", &mutex)

	wg.Wait()
	fmt.Println(msg)
}
