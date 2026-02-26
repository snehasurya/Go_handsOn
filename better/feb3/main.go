package main

import (
	"fmt"
	"sync"
)

type bw struct {
	name string
}

type userDB struct {
	user map[int]string
}

var once sync.Once
var instance *bw
var user *userDB
var mu sync.Mutex
var counter = 1

func GetBWInstance() *bw {
	once.Do(func() {
		instance = &bw{name: "Tibco"}
	})
	return instance
}

func main() {
	s1 := GetBWInstance()
	s2 := GetBWInstance()
	fmt.Println(s1.name)
	fmt.Println(s2.name)
	fmt.Printf("s1 : %p, s2 : %p\n", s1, s2)
	s1.name = "Sneha"
	fmt.Println(s2.name)
	fmt.Printf("s1 : %p, s2 : %p\n", s1, s2)
	for i := 0; i < 1000; i++ {
		go getUserDB(i)
		//go getUserInstace()
	}

	fmt.Scanln()

}

func getUserDB(i int) *userDB {
	if user == nil {
		mu.Lock()
		defer mu.Unlock()
		if user == nil {
			//fmt.Println("single instace of user created")
			fmt.Printf("i got the lock %d\n", i)
			user = &userDB{}
			counter += 1
		} else {
			fmt.Printf("Late Waiter %d\n", i)
			//fmt.Println("user is already created-1, returning that one")
		}
	} else {
		fmt.Printf("Late Arrival %d\n", i)
		//fmt.Println("user is created-2, returning the same ")
	}
	return user
}

func getUserInstace() *userDB {
	once.Do(func() {
		user = &userDB{}
	})
	fmt.Println("user created", user)
	return user
}
