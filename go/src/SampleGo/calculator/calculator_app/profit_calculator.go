package main

import "fmt"

type user struct {
	name   string
	age    int
	gender rune
	salary float64
	fact   bool
}

func main() {

	us := user{
		name:   "sneha",
		age:    30,
		gender: 'F',
	}

	us2 := us
	fmt.Printf("%+v\n", us)
	us2.age = 31
	fmt.Printf("%+v\n", us2)

	fmt.Printf("%+v\n", us)
	us3 := &us
	us3.name = "Vivaan"
	fmt.Printf("us3 %+v\n", us3)
	fmt.Printf("us1  %+v\n", us)

	mp1 := map[string]int64{}
	mp1["apple"] = 1
	mp1["orange"] = 2
	fmt.Println(mp1)
	mp3 := mp1
	mp2 := &mp1
	(*mp2)["orange"] = 3
	fmt.Println("map2", mp2)
	fmt.Println("map1", mp1)
	mp3["apple"] = 4
	fmt.Println("map3", mp3)
	fmt.Println("map1", mp1)

	arr := [4]int{1, 2, 3, 4}
	fmt.Println("array1 ", arr)
	arr[0] = 10
	fmt.Println("array1 ", arr)

	s1 := "sneha"
	fmt.Println("s1 ", s1)
	s1 = "viv"
	fmt.Println("s1 ", s1)
}
