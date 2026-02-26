package main

import "fmt"

type Data struct {
	Value int
}

func (d Data) ModifyByValue() {
	d.Value = 10
}

func (d *Data) ModifyByPointer() {
	d.Value = 20
}

func main() {
	d := Data{Value: 5}
	d.ModifyByValue()
	fmt.Println("After ModifyByValue:", d.Value)
	d.ModifyByPointer()
	fmt.Println("After ModifyByPointer:", d.Value)
	fmt.Println(5 + "A")
	in := []int{}
	fmt.Println(in)
}
