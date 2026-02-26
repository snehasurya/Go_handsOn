package main

import "fmt"

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Square struct {
	side float64
}

func (sq *Square) area() float64 {
	return sq.side * sq.side
}

func main() {
	var s Shape
	c := Circle{
		radius: 10,
	}
	s = &c
	fmt.Println(s.area())
	sq := Square{
		side: 10,
	}
	s = &sq
	fmt.Println(s.area())
}
