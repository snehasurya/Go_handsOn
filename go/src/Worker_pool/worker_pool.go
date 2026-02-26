package main

import (
	"fmt"
	"sync"
	"time"
)

// create a task
type Task interface {
	Process()
}

type User struct {
	UserId   int
	UserName string
}

// process the task
func (u *User) Process() {
	fmt.Printf("the user created is %d with name %s \n", u.UserId, u.UserName)
	time.Sleep(1 * time.Second)
}

type Email struct {
	EmailId string
}

func (e *Email) Process() {
	fmt.Printf("the email created is %s \n", e.EmailId)
	time.Sleep(1 * time.Second)
}

type WorkerPool struct {
	Tasks      []Task
	routinesGo int
	jobChan    chan Task
	wg         sync.WaitGroup
}

func (wp *WorkerPool) Worker() {
	for task := range wp.jobChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.jobChan = make(chan Task)

	for i := 0; i < wp.routinesGo; i++ {
		go wp.Worker()
	}
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.jobChan <- task
	}
	close(wp.jobChan)
	wp.wg.Wait()

}
