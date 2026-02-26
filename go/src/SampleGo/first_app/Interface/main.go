package main

import "fmt"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   int
	height int
}

type Square struct {
	side int
}

func (t Triangle) Area() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}

func (sq *Square) Area() float64 {
	return float64(sq.side) * float64(sq.side)
}

func main() {
	var s Shape

	t := Triangle{
		base:   10,
		height: 20,
	}
	s = t
	fmt.Printf("The area of triangle is %.2f\n", s.Area())

	sq := Square{
		side: 10,
	}
	s = &sq
	fmt.Printf("The area of square is %.2f\n", s.Area())

}
