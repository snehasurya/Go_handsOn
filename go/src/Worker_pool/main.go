package main

import "fmt"

func main() {

	task := []Task{
		&User{UserId: 1, UserName: "Sneha"},
		&Email{EmailId: "SN@gmail.com"},
		&User{UserId: 2, UserName: "Pankaj"},
		&Email{EmailId: "Pa@gmail.com"},
		&User{UserId: 3, UserName: "Vivaan"},
		&Email{EmailId: "viv@gmail.com"},
		&User{UserId: 4, UserName: "Vivaan"},
		&Email{EmailId: "ab@gmail.com"},
		&User{UserId: 5, UserName: "Vivaan"},
		&Email{EmailId: "bcd@gmail.com"},
		&User{UserId: 6, UserName: "Vivaan"},
		&Email{EmailId: "efg@gmail.com"},
		&User{UserId: 7, UserName: "Vivaan"},
		&Email{EmailId: "ahjj@gmail.com"},
		&User{UserId: 8, UserName: "Vivaan"},
		&Email{EmailId: "addf@gmail.com"},
	}

	wp := WorkerPool{
		Tasks:      task,
		routinesGo: 4,
	}
	wp.Run()

	fmt.Println("all work done")
}
