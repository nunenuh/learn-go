package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func main() {
	t := triangle{base: 3, height: 4}
	s := square{side: 5}

	printArea(t)
	printArea(s)
}
