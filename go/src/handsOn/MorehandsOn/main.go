package main

import "fmt"

func main() {
	//maps()
	//stringDemo()
	fmt.Printf("%.3f", 2.22)
}

func maps() {
	map1 := make(map[string]int)
	map2 := make(map[string]int, 1)
	map1["apple"] = 1
	map1["pear"] = 2
	fmt.Println(map1)
	fmt.Println(map2)
	value := map1["apple"]
	fmt.Println(value)
	if _, ok := map1["abc"]; !ok {
		map1["abc"] = 10
	}
	fmt.Println(map1)
	delete(map1, "abc")
	fmt.Println(map1)
}
func slices() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	fmt.Println("S1 : ", s1)
	fmt.Println("S2 : ", s2)
	s2[0] = 99
	fmt.Println("S1 : ", s1)
	fmt.Println("S2 : ", s2)
}

func stringDemo() {
	input := "Sneha 世界"

	for i, v := range input {
		fmt.Printf("at %d : %c  \n", i, v)
	}
	for i := 0; i < len(input); i++ {
		fmt.Printf("%d   %T  \n", input[i], input[i])
	}

}
