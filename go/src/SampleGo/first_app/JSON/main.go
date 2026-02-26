package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name   string `json: "Name"`
	Age    int    `json: "Age"`
	Gender rune   `json: "-"`
}

func main() {

	p := person{
		Name:   "Sneha",
		Age:    30,
		Gender: 'F',
	}
	jsonData, err := json.Marshal(p)
	fmt.Println(jsonData)

	//jsonString := `{"Name" : "Sneha", "Age" : 30, "Gender": "F"}`

	var P2 person
	err = json.Unmarshal(jsonData, &P2)
	if err != nil {
		fmt.Println("error in unmarshal ", err)
	}
	fmt.Printf("Unmarshelled structure %+v\n", P2)
	fmt.Printf("Gender: %c\n", P2.Gender)
}
